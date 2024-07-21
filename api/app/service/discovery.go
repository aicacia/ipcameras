package service

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"time"

	"github.com/aicacia/go-cmap"
	"github.com/aicacia/ipcameras/api/app/config"
	"github.com/aicacia/ipcameras/api/app/util"
	"github.com/use-go/onvif"
	onvifdevice "github.com/use-go/onvif/device"
	"github.com/use-go/onvif/media"
	devicesdk "github.com/use-go/onvif/sdk/device"
	mediasdk "github.com/use-go/onvif/sdk/media"
	xsdonvif "github.com/use-go/onvif/xsd/onvif"
)

type discoveryDeviceST struct {
	device onvif.Device
	info   onvif.DeviceInfo
}

type DeviceST struct {
	Device    onvif.Device
	Info      onvif.DeviceInfo
	MediaUris map[string]xsdonvif.MediaUri
	LastSeen  time.Time
}

func (d *DeviceST) HardwareId() string {
	return d.Info.HardwareId
}

func (d *DeviceST) Name() string {
	manufacturer := d.Info.Manufacturer
	model := d.Info.Model
	if manufacturer != "" && model != "" {
		return fmt.Sprintf("%s %s", manufacturer, model)
	}
	if manufacturer != "" {
		return manufacturer
	}
	if model != "" {
		return model
	}
	return d.Info.HardwareId
}

var Devices = cmap.New[string, *DeviceST]()

func InitDiscovery() {
	ctx := context.Background()
	go discovery(ctx)
}

func discovery(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("discovery panic", "error", err)
			go discovery(ctx)
		}
	}()
	for {
		if !config.Get().Discovery.Enabled {
			time.Sleep(time.Duration(time.Minute))
			continue
		}
		interfaceNames, err := localInterfaceNames()
		if err != nil {
			slog.Error("error getting interfaces", "error", err)
		}
		discoveredDevices := make(map[string]discoveryDeviceST, 0)
		for _, interfaceName := range interfaceNames {
			interfaceDevices, err := onvif.GetAvailableDevicesAtSpecificEthernetInterface(interfaceName)
			if err != nil {
				slog.Error("error getting devices", "error", err)
			}
			if len(interfaceDevices) == 0 {
				continue
			}
			for _, device := range interfaceDevices {
				info := device.GetDeviceInfo()
				if info.HardwareId == "" {
					deviceInfoResponse, err := devicesdk.Call_GetDeviceInformation(ctx, &device, onvifdevice.GetDeviceInformation{})
					if err != nil {
						slog.Error("error getting Device information", "error", err)
						continue
					}
					info.Manufacturer = deviceInfoResponse.Manufacturer
					info.Model = deviceInfoResponse.Model
					info.FirmwareVersion = deviceInfoResponse.FirmwareVersion
					info.SerialNumber = deviceInfoResponse.SerialNumber
					info.HardwareId = deviceInfoResponse.HardwareId

				}
				discoveredDevices[info.HardwareId] = discoveryDeviceST{device, info}
			}
		}
		for _, discoveredDevice := range discoveredDevices {
			device, ok := Devices.Get(discoveredDevice.info.HardwareId)
			if !ok {
				device = &DeviceST{Device: discoveredDevice.device, Info: discoveredDevice.info, MediaUris: map[string]xsdonvif.MediaUri{}}
				profiles, err := mediasdk.Call_GetProfiles(ctx, &discoveredDevice.device, media.GetProfiles{})
				if err != nil {
					slog.Error("error getting device profiles", "error", err)
					continue
				}
				for _, profile := range profiles.Profiles {
					streamUri, err := mediasdk.Call_GetStreamUri(ctx, &discoveredDevice.device, media.GetStreamUri{ProfileToken: profile.Token})
					if err != nil {
						slog.Error("error getting device profile stream uri", "error", err)
						continue
					}
					name := util.GetUniqueKey(device.MediaUris, string(profile.Name))
					device.MediaUris[name] = streamUri.MediaUri
				}
				Devices.Set(device.HardwareId(), device)
				slog.Info("discovered device", "hardwareId", discoveredDevice.info.HardwareId)
				camera, err := GetCameraByHardwareId(device.HardwareId())
				if err != nil {
					slog.Error("error getting camera", "error", err)
					continue
				}
				onAddCamera(cameraFromDevice(camera, device))
			} else {
				slog.Debug("device already exists", "hardwareId", discoveredDevice.info.HardwareId)
			}
			device.LastSeen = time.Now()
		}
		for entry := range Devices.Iter() {
			if _, ok := discoveredDevices[entry.Key]; !ok {
				if time.Since(entry.Val.LastSeen) >= time.Duration(config.Get().Discovery.RemoveTimeoutMilliseconds) {
					Devices.Remove(entry.Key)
					slog.Debug("removed device", "hardwareId", entry.Key)
				}
			}
		}
	}
}

func localInterfaceNames() ([]string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var interfaceNames []string
	for _, i := range interfaces {
		interfaceNames = append(interfaceNames, i.Name)
	}
	return interfaceNames, nil
}
