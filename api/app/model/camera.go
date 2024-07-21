package model

import (
	"time"

	"github.com/aicacia/ipcameras/api/app/service"
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

func CameraFromService(c *service.CameraST) CameraST {
	var recordWindow *int64
	if c.RecordWindow != nil {
		value := int64(*c.RecordWindow)
		recordWindow = &value
	}
	return CameraST{
		HardwareId:   c.HardwareId,
		Name:         c.Name,
		MediaUris:    c.MediaUris,
		Saved:        c.Saved,
		Record:       c.Record,
		RecordWindow: recordWindow,
		Discovered:   c.Discovered,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
	}
}

type UpsertCameraST struct {
	service.UpsertCameraST
} // @name UpsertCamera
