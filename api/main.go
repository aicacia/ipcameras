package main

import (
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/aicacia/ipcameras/api/app"
	"github.com/aicacia/ipcameras/api/app/config"
	"github.com/aicacia/ipcameras/api/app/router"
	arg "github.com/alexflint/go-arg"
)

var (
	Version string = "0.1.0"
	Build   string = fmt.Sprint(time.Now().UnixMilli())
)

var args struct {
	ConfigPath string `arg:"-c,--config" help:"path to config"`
}

// @title IPCameras API
// @description IPCameras API API
// @contact.name Nathan Faucett
// @contact.email nathanfaucett@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
// @securityDefinitions.apikey Locale
// @in header
// @name X-Locale
// @securityDefinitions.apikey Timezone
// @in header
// @name X-Timezone
func main() {
	defer func() {
		rec := recover()
		if rec != nil {
			log.Fatalf("application panic: %v\n", rec)
		}
	}()
	arg.MustParse(&args)
	fiberApp := app.InitApp(app.AppConfigST{
		ConfigPath: args.ConfigPath,
		Version:    Version,
		Build:      Build,
	})
	router.InstallRouter(fiberApp)

	addr := fmt.Sprintf("%s:%d", config.Get().Host, config.Get().Port)
	slog.Info("Listening on %s\n", "addr", addr)

	log.Fatal(fiberApp.Listen(addr))
}
