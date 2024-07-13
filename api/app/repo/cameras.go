package repo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/aicacia/ipcameras/api/app/config"
	"github.com/aicacia/ipcameras/api/app/service"
	"github.com/aicacia/ipcameras/api/app/util"
)

type CameraST struct {
	HardwareId   string            `json:"hardwareId" validate:"required"`
	Name         string            `json:"name" validate:"required"`
	MediaUris    map[string]string `json:"mediaUris" validate:"required"`
	Discovered   bool              `json:"discovered" validate:"required"`
	Saved        bool              `json:"saved" validate:"required"`
	Record       bool              `json:"record" validate:"required"`
	RecordWindow *time.Duration    `json:"recordWindow" validate:"required"`
	CreatedAt    *time.Time        `json:"createdAt" format:"date-time"`
	UpdatedAt    time.Time         `json:"updatedAt" validate:"required" format:"date-time"`
	Device       *service.DeviceST `json:"-"`
}

func GetCameras() ([]*CameraST, error) {
	var errs []error
	cameras, err := readCameras()
	if err != nil {
		errs = append(errs, err)
	}
	for device := range service.Devices.Values() {
		camera, ok := cameras[device.HardwareId()]
		if !ok {
			camera = &CameraST{
				HardwareId: device.HardwareId(),
				Name:       device.Name(),
				MediaUris:  map[string]string{},
			}
			cameras[camera.HardwareId] = camera
		}
		camera.Device = device
		for name, mediaUri := range device.MediaUris {
			camera.MediaUris[name] = string(mediaUri.Uri)
		}
		camera.UpdatedAt = device.LastSeen
		camera.Discovered = true
	}
	return util.Values(cameras), errors.Join(errs...)
}

func GetCameraByHardwareId(hardwareId string) (*CameraST, error) {
	camera, cameraDiskErr := readCamera(hardwareId)
	device, deviceExists := service.Devices.Get(hardwareId)
	if deviceExists {
		if camera == nil {
			camera = &CameraST{
				MediaUris: map[string]string{},
			}
		}
		camera.HardwareId = device.HardwareId()
		camera.Name = device.Name()
		camera.Device = device
		for name, mediaUri := range device.MediaUris {
			camera.MediaUris[name] = string(mediaUri.Uri)
		}
		camera.UpdatedAt = device.LastSeen
		camera.Discovered = true
	}
	return camera, cameraDiskErr
}

type UpsertCameraST struct {
	HardwareId   string            `json:"hardwareId" validate:"required"`
	Name         string            `json:"name" validate:"required"`
	MediaUris    map[string]string `json:"mediaUris" validate:"required"`
	Record       bool              `json:"record" validate:"required"`
	RecordWindow *time.Duration    `json:"recordWindow"`
}

func UpsertCameraByHardwareId(hardwareId string, update UpsertCameraST) (*CameraST, error) {
	camera, _ := GetCameraByHardwareId(update.HardwareId)
	if camera == nil {
		camera = &CameraST{
			MediaUris: map[string]string{},
		}
	}
	for name, mediaUri := range update.MediaUris {
		camera.MediaUris[name] = mediaUri
	}
	camera.HardwareId = hardwareId
	if update.Name != "" {
		camera.Name = update.Name
	}
	if camera.CreatedAt == nil {
		now := time.Now()
		camera.CreatedAt = &now
	}
	camera.Record = update.Record
	camera.RecordWindow = update.RecordWindow
	camera.Saved = true
	if err := writeCamera(camera); err != nil {
		return nil, err
	}
	return camera, nil
}

func readCameras() (map[string]*CameraST, error) {
	entries, err := os.ReadDir(config.Get().Cameras.Path)
	if err != nil {
		return nil, err
	}
	cameras := make(map[string]*CameraST, len(entries))
	var errs []error
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		ext := path.Ext(entry.Name())
		if ext != ".json" {
			continue
		}
		hardwareId := strings.TrimSuffix(entry.Name(), ext)
		camera, err := readCamera(hardwareId)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		if camera == nil {
			continue
		}
		cameras[camera.HardwareId] = camera
	}
	return cameras, errors.Join(errs...)
}

func readCamera(hardwareId string) (*CameraST, error) {
	var camera CameraST
	bytes, err := os.ReadFile(path.Join(config.Get().Cameras.Path, fmt.Sprintf("%s.json", hardwareId)))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &camera)
	if err != nil {
		return nil, err
	}
	return &camera, nil
}

func writeCamera(camera *CameraST) error {
	err := os.MkdirAll(config.Get().Cameras.Path, os.ModePerm)
	if err != nil {
		return err
	}
	cameraBytes, err := json.MarshalIndent(camera, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(path.Join(config.Get().Cameras.Path, fmt.Sprintf("%s.json", camera.HardwareId)), cameraBytes, os.ModePerm)
}
