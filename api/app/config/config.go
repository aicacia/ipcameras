package config

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	atomic_value "github.com/aicacia/go-atomic-value"
	"github.com/aicacia/ipcameras/api/app/util"
)

var config atomic_value.AtomicValue[*ConfigST]

func Get() *ConfigST {
	return config.Load()
}

type DashboardST struct {
	Enabled bool `json:"enabled"`
}

type OpenAPIST struct {
	Enabled bool `json:"enabled"`
}

type CamerasST struct {
	Path string `json:"path"`
}

type UsersST struct {
	Path string `json:"path"`
}

type DiscoveryST struct {
	Enabled                   bool  `json:"enabled"`
	RemoveTimeoutMilliseconds int64 `json:"removeTimeoutMilliseconds"`
}

type P2PST struct {
	Id        string `json:"id"`
	Password  string `json:"password"`
	Host      string `json:"host"`
	SSL       bool   `json:"ssl"`
	JWTSecret string `json:"jwtSecret"`
}

type IceST struct {
	Servers    []string `json:"servers"`
	Username   string   `json:"username"`
	Credential string   `json:"credential"`
}

type RTSPST struct {
	ConnectTimeoutSeconds int  `json:"connectTimeoutSeconds"`
	IOTimeoutSeconds      int  `json:"ioTimeoutSeconds"`
	Debug                 bool `json:"debug"`
}

type JWTST struct {
	Secret                  string `json:"secret"`
	ExpiresInSeconds        int64  `json:"expiresInSeconds"`
	RefreshExpiresInSeconds int64  `json:"refreshExpiresInSeconds"`
}

type ConfigST struct {
	Host      string      `json:"host"`
	Port      int16       `json:"port"`
	URL       string      `json:"url"`
	Dashboard DashboardST `json:"dashboard"`
	OpenAPI   OpenAPIST   `json:"openapi"`
	Cameras   CamerasST   `json:"cameras"`
	Users     UsersST     `json:"users"`
	Discovery DiscoveryST `json:"discovery"`
	P2P       P2PST       `json:"p2p"`
	Ice       IceST       `json:"ice"`
	RTSP      RTSPST      `json:"rtsp"`
	JWT       JWTST       `json:"jwt"`
}

func (config *ConfigST) GetP2PWSURL() string {
	protocal := "ws"
	if config.P2P.SSL {
		protocal = "wss"
	}
	return fmt.Sprintf("%s://%s", protocal, config.P2P.Host)
}

func (config *ConfigST) GetP2PURL() string {
	protocal := "http"
	if config.P2P.SSL {
		protocal = "https"
	}
	return fmt.Sprintf("%s://%s", protocal, config.P2P.Host)
}

func InitConfig(path string) error {
	if path == "" {
		path = "./config.json"
	}
	configBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	c := ConfigST{
		Host:      "0.0.0.0",
		Port:      3000,
		URL:       "http://localhost:3000",
		Dashboard: DashboardST{Enabled: true},
		OpenAPI: OpenAPIST{
			Enabled: true,
		},
		Cameras: CamerasST{
			Path: "./data/cameras",
		},
		Users: UsersST{
			Path: "./data/users",
		},
		Ice: IceST{
			Servers: []string{"stun:stun.l.google.com:19302"},
		},
		Discovery: DiscoveryST{Enabled: true, RemoveTimeoutMilliseconds: 60000},
		P2P: P2PST{
			Host: "localhost:4000",
			SSL:  false,
		},
		RTSP: RTSPST{
			ConnectTimeoutSeconds: 10,
			IOTimeoutSeconds:      10,
			Debug:                 true,
		},
		JWT: JWTST{
			ExpiresInSeconds:        86400,
			RefreshExpiresInSeconds: 604800,
		},
	}
	err = json.Unmarshal(configBytes, &c)
	if err != nil {
		return err
	}
	if c.JWT.Secret == "" {
		bytes, err := util.GenerateRandomBytes(256)
		if err != nil {
			slog.Error("failed to generate random bytes for JWT secret", "error", err)
		}
		c.JWT.Secret = base64.StdEncoding.EncodeToString(bytes)
	}
	config.Store(&c)
	return nil
}
