package controller

import (
	"net/http"
	"time"

	"github.com/aicacia/ipcameras/api/app"
	"github.com/aicacia/ipcameras/api/app/config"
	"github.com/aicacia/ipcameras/api/app/model"
	"github.com/aicacia/ipcameras/api/docs"
	"github.com/gofiber/fiber/v2"
)

func GetOpenAPI(c *fiber.Ctx) error {
	c.Status(http.StatusOK)
	c.Set("Content-Type", "application/json; charset=utf-8")
	return c.SendString(docs.SwaggerInfo.ReadDoc())
}

// GetHealthCheck
//
//	@ID		  		health-check
//	@Summary		Get Health Check
//	@Tags		  	app
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.HealthST
//	@Failure		500	{object}	model.HealthST
//	@Router			/health [get]
func GetHealthCheck(c *fiber.Ctx) error {
	health := model.HealthST{
		Date: time.Now().UTC(),
	}
	if health.IsHealthy() {
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusInternalServerError)
	}
	return c.JSON(health)
}

// GetVersion
//
//	@ID			  	version
//	@Summary		Get Version
//	@Tags			  app
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.VersionST
//	@Router			/version [get]
func GetVersion(c *fiber.Ctx) error {
	c.Status(http.StatusOK)
	return c.JSON(app.Version)
}

// GetP2P
//
//	@ID			  	p2p-access
//	@Summary		Get p2p access info
//	@Tags			  app
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.P2PAccessST
//	@Router			/p2p-access [get]
//
//	@Security		Authorization
func GetP2PAccess(c *fiber.Ctx) error {
	c.Status(http.StatusOK)
	return c.JSON(model.P2PAccessST{
		Host:     config.Get().P2P.Host,
		SSL:      config.Get().P2P.SSL,
		Id:       config.Get().P2P.Id,
		Password: config.Get().P2P.Password,
	})
}
