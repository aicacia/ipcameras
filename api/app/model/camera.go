package model

import (
	"time"

	"github.com/aicacia/ipcameras/api/app/repo"
)

type CameraST struct {
	HardwareId   string            `json:"hardwareId" validate:"required"`
	Name         string            `json:"name" validate:"required"`
	MediaUris    map[string]string `json:"mediaUris" validate:"required"`
	Saved        bool              `json:"saved" validate:"required"`
	Record       bool              `json:"record" validate:"required"`
	RecordWindow *int64            `json:"recordWindow"`
	Discovered   bool              `json:"discovered" validate:"required"`
	CreatedAt    *time.Time        `json:"createdAt" format:"date-time"`
	UpdatedAt    time.Time         `json:"updatedAt" validate:"required" format:"date-time"`
} // @name Camera

func CameraFromRepo(repo *repo.CameraST) CameraST {
	var recordWindow *int64
	if repo.RecordWindow != nil {
		value := int64(*repo.RecordWindow)
		recordWindow = &value
	}
	return CameraST{
		HardwareId:   repo.HardwareId,
		Name:         repo.Name,
		MediaUris:    repo.MediaUris,
		Saved:        repo.Saved,
		Record:       repo.Record,
		RecordWindow: recordWindow,
		Discovered:   repo.Discovered,
		CreatedAt:    repo.CreatedAt,
		UpdatedAt:    repo.UpdatedAt,
	}
}

type UpsertCameraST struct {
	repo.UpsertCameraST
} // @name UpsertCamera
