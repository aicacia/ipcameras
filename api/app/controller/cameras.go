package controller

import (
	"log/slog"
	"net/http"

	"github.com/aicacia/ipcameras/api/app/model"
	"github.com/aicacia/ipcameras/api/app/repo"
	"github.com/aicacia/ipcameras/api/app/util"
	"github.com/gofiber/fiber/v2"
)

// GetCameras
//
//	@ID				cameras
//	@Summary		Get all cameras
//	@Tags			camera
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	model.CameraST
//	@Router			/cameras [get]
func GetCameras(c *fiber.Ctx) error {
	cameras, err := repo.GetCameras()
	if err != nil {
		slog.Error("error getting cameras", "error")
	}
	return c.JSON(util.Map(cameras, model.CameraFromRepo))
}

// GetCameraByHardwareId
//
//	@ID				cameraByHardwareId
//	@Summary		Get camera by hardware id
//	@Tags			camera
//	@Accept			json
//	@Produce		json
//	@Param			hardwareId	path		string	true	"Hardware Id"
//	@Success		200	{object}	model.CameraST
//	@Failure		404	{object}	model.ErrorST
//	@Router			/cameras/{hardwareId} [get]
func GetCameraByHardwareId(c *fiber.Ctx) error {
	camera, err := repo.GetCameraByHardwareId(c.Params("hardwareId"))
	if err != nil {
		slog.Error("error getting camera", "error", err)
	}
	if camera == nil {
		return model.NewError(http.StatusNotFound).AddError("hardwareId", "invalid")
	}
	return c.JSON(model.CameraFromRepo(camera))
}

// PatchCameraByHardwareId
//
//	@ID				  updateCameraByHardwareId
//	@Summary		update camera by hardware id
//	@Tags			  camera
//	@Accept			json
//	@Produce		json
//	@Param			hardwareId	path		string	true	"Hardware Id"
//	@Param      updates body     model.UpsertCameraST true "Camera"
//	@Success		200	{object}	model.CameraST
//	@Failure		404	{object}	model.ErrorST
//	@Router			/cameras/{hardwareId} [patch]
func PatchCameraByHardwareId(c *fiber.Ctx) error {
	var body model.UpsertCameraST
	if err := c.BodyParser(&body); err != nil {
		slog.Error("error parsing body", "error", err)
		return model.NewError(http.StatusBadRequest).AddError("updates", "invalid")
	}
	updatedCamera, err := repo.UpsertCameraByHardwareId(c.Params("hardwareId"), body.UpsertCameraST)
	if err != nil {
		slog.Error("error updating camera", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	if updatedCamera == nil {
		return model.NewError(http.StatusNotFound).AddError("hardwareId", "invalid")
	}
	return c.JSON(model.CameraFromRepo(updatedCamera))
}
