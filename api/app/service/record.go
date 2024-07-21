package service

import (
	"errors"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/aicacia/go-cmap"
	"github.com/aicacia/ipcameras/api/app/config"
	"github.com/aicacia/ipcameras/api/app/service/format"
	"github.com/aicacia/ipcameras/api/app/util"
	"github.com/deepch/vdk/av"
)

var ErrInvalidRecording = errors.New("invalid recording")

type recorderST struct {
	viewers cmap.CMap[string, *ViewerST]
}

var recorders = cmap.New[string, *recorderST]()

func InitRecord() {
	go record()
}

func record() {
	subscriber := CameraEventPubSub.Subscribe()

	defer subscriber.Close()
	defer func() {
		if err := recover(); err != nil {
			slog.Error("record panic", "error", err)
			go record()
		}
	}()

	for e := range subscriber.C {
		switch e.Type() {
		case CameraAdded:
			event := e.(*AddCameraEvent)
			if event.Camera.Record {
				for streamName := range event.Camera.MediaUris {
					addRecorder(event.Camera, streamName)
				}
			}
		case CameraUpdated:
			event := e.(*UpdateCameraEvent)
			if event.Camera.Record && event.PrevCamera != nil && !event.PrevCamera.Record {
				slog.Debug("camera set to record", "camera", event.Camera.HardwareId)
			} else {
				slog.Debug("stop recording", "camera", event.Camera.HardwareId)
			}
		case CameraDeleted:
			event := e.(*DeleteCameraEvent)
			slog.Debug("stop recording", "camera", event.Camera.HardwareId)
		}
	}
}

func addRecorder(camera *CameraST, streamName string) {
	streamUri, ok := camera.MediaUris[streamName]
	if !ok {
		slog.Info("stream not found", "camera", camera.HardwareId, "stream", streamName)
		return
	}
	recorder, ok := recorders.Get(camera.HardwareId)
	if !ok || recorder == nil {
		recorder = &recorderST{
			viewers: cmap.New[string, *ViewerST](),
		}
		recorders.Set(camera.HardwareId, recorder)
	}
	if err := RunIfNotRunning(streamUri); err != nil {
		slog.Error("failed to run stream", "camera", camera.HardwareId, "stream", streamName, "error", err)
		return
	}
	viewer, ok := recorder.viewers.Get(streamUri)
	if !ok || viewer == nil {
		viewer = AddViewer(streamUri, streamName)
		recorder.viewers.Set(streamName, viewer)
	}
	if viewer != nil {
		go startRecording(camera.HardwareId, streamName, streamUri, viewer.Socket)
	}
}

func GetRecordingFolderPathForTime(cameraHardwareId, streamName string, t *time.Time) string {
	return path.Join(
		config.Get().Recording.Folder,
		cameraHardwareId,
		streamName,
		strconv.Itoa(t.Year()),
		strconv.Itoa(int(t.Month())),
		strconv.Itoa(t.Day()),
		strconv.Itoa(t.Hour()),
		strconv.Itoa(t.Minute()),
	)
}

func recordingPathToTime(cameraHardwareId, streamName, recordingPath string) (*time.Time, error) {
	dateStr := recordingPath[len(path.Join(
		config.Get().Recording.Folder,
		cameraHardwareId,
		streamName))+1:]
	parts := strings.SplitN(dateStr, string(os.PathSeparator), 5)
	year, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}
	month, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	day, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, err
	}
	hour, err := strconv.Atoi(parts[3])
	if err != nil {
		return nil, err
	}
	minute, err := strconv.Atoi(parts[4])
	if err != nil {
		return nil, err
	}
	time := time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)
	return &time, nil
}

func GetCameraTimesOnDisk(cameraHardwareId, streamName string) ([]*time.Time, error) {
	pattern := path.Join(
		config.Get().Recording.Folder,
		cameraHardwareId,
		streamName,
		"*",
		"*",
		"*",
		"*",
		"*",
	)
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}
	if len(matches) < 1 {
		return nil, ErrInvalidRecording
	}
	dates := make([]*time.Time, 0, len(matches))
	for _, match := range matches {
		t, err := recordingPathToTime(cameraHardwareId, streamName, match)
		if err != nil {
			return nil, err
		}
		dates = append(dates, t)
	}
	sort.Slice(dates, func(i, j int) bool {
		return dates[i].Before(*dates[j])
	})
	return dates, nil
}

func GetCameraRecordingRanges(cameraHardwareId, streamName string) ([]util.TimeRangeST, error) {
	pattern := path.Join(
		config.Get().Recording.Folder,
		cameraHardwareId,
		"*",
		"*",
		"*",
		"*",
		"*",
	)
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}
	if len(matches) < 1 {
		return nil, ErrInvalidRecording
	}
	dates := sort.StringSlice(matches)
	var times []*time.Time
	for _, date := range dates {
		t, err := recordingPathToTime(cameraHardwareId, streamName, date)
		if err == nil {
			times = append(times, t)
		}
	}
	ranges := make([]util.TimeRangeST, 0)
	if len(times) == 0 {
		return ranges, nil
	}
	startTime := times[0]
	for i := 0; i < len(times)-1; i++ {
		prev := times[i]
		next := times[i+1]
		if prev.Minute() != next.Minute()-1 {
			ranges = append(ranges, util.TimeRangeST{
				Start: *startTime,
				End:   prev,
			})
			startTime = next
		}
	}
	ranges = append(ranges, util.TimeRangeST{
		Start: *startTime,
		End:   nil,
	})
	return ranges, nil
}

func startRecording(cameraHardwareId, streamName, url string, packets chan *av.Packet) {
	var muxer *format.Muxer
	var nextMinute time.Time
	for packet := range packets {
		currentTime := format.GetPacketTime(packet)
	Muxer:
		if muxer == nil {
			var err error
			muxer, err = format.NewMuxer(
				GetRecordingFolderPathForTime(cameraHardwareId, streamName, &currentTime),
			)
			if err != nil {
				slog.Error("Failed to create raw muxer", "url", url, "error", err)
				return
			}
			err = muxer.WriteHeader(GetClientCurrentCodecs(url))
			if err != nil {
				muxer.Close()
				slog.Error("Failed to write codecs", "url", url, "error", err)
				return
			}
			nextMinute = util.TruncateToMinute(currentTime.Add(time.Minute))
		}
		if currentTime.After(nextMinute) {
			muxer.Close()
			muxer = nil
			goto Muxer
		}
		if err := muxer.WritePacket(packet); err != nil {
			slog.Error("Failed to write packet", "url", url, "error", err)
		}
	}
}
