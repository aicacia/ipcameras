package app

import (
	"errors"
	"log/slog"
	"net/http"
	"net/url"
	"os"

	"github.com/aicacia/ipcameras/api/app/config"
	"github.com/aicacia/ipcameras/api/app/model"
	"github.com/aicacia/ipcameras/api/app/service"
	"github.com/aicacia/ipcameras/api/docs"
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

var Version model.VersionST

type AppConfigST struct {
	ConfigPath string
	Build      string
	Version    string
}

func InitApp(appConfig AppConfigST) *fiber.App {
	var envs []string
	env := os.Getenv("APP_ENV")
	if env != "" {
		envs = append(envs, ".env."+env)
	}
	envs = append(envs, ".env")
	if err := godotenv.Load(envs...); err != nil {
		slog.Error("error loading .env file", "error", err)
	}
	if err := config.InitConfig(appConfig.ConfigPath); err != nil {
		slog.Error("error initializing config", "error", err)
		return nil
	}

	loggerWriter := os.Stdout
	logger := slog.New(slog.NewTextHandler(loggerWriter, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	if err := service.InitUsers(); err != nil {
		slog.Error("error initializing users", "error", err)
		return nil
	}
	service.InitDiscovery()
	service.InitRecord()

	Version.Version = appConfig.Version
	Version.Build = appConfig.Build

	docs.SwaggerInfo.Version = appConfig.Version
	uri, err := url.Parse(config.Get().URL)
	if err != nil {
		slog.Error("error parsing URI", "error", err)
		return nil
	}
	docs.SwaggerInfo.Host = uri.Host

	// https://docs.gofiber.io/api/fiber#config
	fiberApp := fiber.New(fiber.Config{
		Prefork:       false,
		Network:       "",
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "",
		AppName:       "",
		ErrorHandler:  errorHandler,
	})
	fiberApp.Use(fiberRecover.New())
	fiberApp.Use(fiberLogger.New(fiberLogger.Config{
		Output:     loggerWriter,
		TimeZone:   "UTC",
		TimeFormat: "2006/01/02 15:04:05",
		Format:     "${time} ${status} - ${ip} ${latency} ${method} ${path}\n",
	}))
	if config.Get().Dashboard.Enabled {
		fiberApp.Use("/dashboard", monitor.New())
	}
	service.InitPeerListener()
	service.InitRTSP()
	service.InitWebRTCServer(fiberApp)

	return fiberApp
}

func errorHandler(c *fiber.Ctx, err error) error {
	if e, ok := err.(*model.ErrorST); ok {
		return e.Send(c)
	}
	code := http.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	slog.Debug("error", "code", code, "err", err)
	return model.NewError(code).AddError("internal", err.Error()).Send(c)
}
