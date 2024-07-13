package app

import (
	"errors"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"os"

	"github.com/aicacia/ipcameras/api/app/config"
	"github.com/aicacia/ipcameras/api/app/model"
	"github.com/aicacia/ipcameras/api/app/repo"
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
	err := godotenv.Load(envs...)
	if err != nil {
		log.Printf("error loading .env file: %v\n", err)
	}
	err = config.InitConfig(appConfig.ConfigPath)
	if err != nil {
		log.Fatalf("error initializing config: %v\n", err)
	}

	loggerWriter := os.Stdout
	logger := slog.New(slog.NewTextHandler(loggerWriter, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	err = service.InitDiscovery()
	if err != nil {
		log.Fatalf("error initializing discovery: %v\n", err)
	}
	err = repo.InitUsers()
	if err != nil {
		log.Fatalf("error initializing users: %v\n", err)
	}

	Version.Version = appConfig.Version
	Version.Build = appConfig.Build

	docs.SwaggerInfo.Version = appConfig.Version
	uri, err := url.Parse(config.Get().URL)
	if err != nil {
		log.Fatalf("error parsing URI: %v\n", err)
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
		ErrorHandler:  ErrorHandler,
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
	err = service.InitPeerListener()
	if err != nil {
		log.Fatalf("error initializing Peer Listener: %v\n", err)
	}
	service.InitRTSP()
	service.InitWebRTCServer(fiberApp)

	return fiberApp
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	if e, ok := err.(*model.ErrorST); ok {
		return e.Send(c)
	}
	code := http.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	return model.NewError(code).AddError("internal", err.Error()).Send(c)
}
