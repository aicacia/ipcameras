package service

import (
	"github.com/aicacia/go-pubsub"
)

const (
	CameraAdded = 1 + iota
	CameraUpdated
	CameraDeleted
)

type CameraEvent interface {
	Type() int
}

type AddCameraEvent struct {
	Camera *CameraST
}

func (e *AddCameraEvent) Type() int {
	return CameraAdded
}

type UpdateCameraEvent struct {
	Camera     *CameraST
	PrevCamera *CameraST
}

func (e *UpdateCameraEvent) Type() int {
	return CameraUpdated
}

type DeleteCameraEvent struct {
	Camera *CameraST
}

func (e *DeleteCameraEvent) Type() int {
	return CameraDeleted
}

var CameraEventPubSub = pubsub.NewPubSub[CameraEvent]()

func onAddCamera(camera *CameraST) {
	var event CameraEvent = &AddCameraEvent{
		Camera: camera,
	}
	CameraEventPubSub.Publish(event)
}

func onUpdateCamera(camera, prevCamera *CameraST) {
	var event CameraEvent = &UpdateCameraEvent{
		Camera:     camera,
		PrevCamera: prevCamera,
	}
	CameraEventPubSub.Publish(event)
}

func onDeleteCamera(camera *CameraST) {
	var event CameraEvent = &DeleteCameraEvent{
		Camera: camera,
	}
	CameraEventPubSub.Publish(event)
}
