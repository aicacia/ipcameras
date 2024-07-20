package router

import (
	"github.com/aicacia/ipcameras/api/app/config"
	"github.com/aicacia/ipcameras/api/app/controller"
	"github.com/aicacia/ipcameras/api/app/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func InstallRouter(fiberApp *fiber.App) {
	root := fiberApp.Group("",
		cors.New(cors.Config{
			AllowOriginsFunc: func(origin string) bool {
				return true
			},
			AllowCredentials: true,
		}),
		middleware.LocaleMiddleware(),
		middleware.TimezoneMiddleware(),
	)

	if config.Get().OpenAPI.Enabled {
		root.Get("/openapi.json", controller.GetOpenAPI)
	}

	root.Get("/health", controller.GetHealthCheck)
	root.Get("/version", controller.GetVersion)
	root.Get("/ice-servers", controller.GetICEServers)

	root.Post("/token", controller.PostToken)

	auth := root.Group("", middleware.AuthorizedMiddleware())
	auth.Get("/p2p-access", controller.GetP2PAccess)

	user := auth.Group("/user")
	user.Get("", controller.GetCurrentUser)
	user.Patch("/reset-password", controller.PatchResetPassword)

	cameras := auth.Group("/cameras")
	cameras.Get("", controller.GetCameras)
	cameraByHardwareId := cameras.Group("/:hardwareId")
	cameraByHardwareId.Get("", controller.GetCameraByHardwareId)
	cameraByHardwareId.Patch("", controller.PatchCameraByHardwareId)
}
