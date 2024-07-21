package model

import (
	"time"

	"github.com/aicacia/ipcameras/api/app/service"
)

type UserST struct {
	Username  string    `json:"username" validate:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required" format:"date-time"`
	UpdatedAt time.Time `json:"updatedAt" validate:"required" format:"date-time"`
} // @name User

func UserFromService(u *service.UserST) UserST {
	return UserST{
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

type ResetPasswordST struct {
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required"`
} // @name ResetPassword
