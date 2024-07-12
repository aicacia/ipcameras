package router

import (
	"github.com/aicacia/ipcameras/api/app/config"
	"github.com/aicacia/ipcameras/api/app/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func InstallRouter(fiberApp *fiber.App) {
	root := fiberApp.Group("", cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return true
		},
		AllowCredentials: true,
	}))

	if config.Get().OpenAPI.Enabled {
		root.Get("/openapi.json", controller.GetOpenAPI)
	}

	root.Get("/health", controller.GetHealthCheck)
	root.Get("/version", controller.GetVersion)
	root.Get("/p2p-access", controller.GetP2PAccess)

	cameras := root.Group("/cameras")
	cameras.Get("", controller.GetCameras)
	cameraByHardwareId := cameras.Group("/:hardwareId")
	cameraByHardwareId.Get("", controller.GetCameraByHardwareId)
	cameraByHardwareId.Patch("", controller.PatchCameraByHardwareId)
}
