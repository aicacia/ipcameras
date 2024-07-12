package model

import "time"

type HealthST struct {
	Date time.Time `json:"date" validate:"required" format:"date-time"`
} // @name Health

func (health HealthST) IsHealthy() bool {
	return true
}

type VersionST struct {
	Version string `json:"version" validate:"required"`
	Build   string `json:"build" validate:"required"`
} // @name Version

type OffsetAndLimitQueryST struct {
	Offset *int `query:"offset"`
	Limit  *int `query:"limit"`
} // @name OffsetAndLimitQuery

type P2PAccessST struct {
	Host     string `json:"host" validate:"required"`
	SSL      bool   `json:"ssl" validate:"required"`
	Id       string `json:"id" validate:"required"`
	Password string `json:"password" validate:"required"`
} // @name P2PAccess
