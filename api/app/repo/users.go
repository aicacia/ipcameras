package repo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/aicacia/ipcameras/api/app/config"
	"github.com/aicacia/ipcameras/api/app/util"
)

type UserST struct {
	Username          string    `json:"username" validate:"required"`
	EncryptedPassword string    `json:"encryptedPassword" validate:"required"`
	CreatedAt         time.Time `json:"createdAt" validate:"required" format:"date-time"`
	UpdatedAt         time.Time `json:"updatedAt" validate:"required" format:"date-time"`
}

func GetUsers() ([]*UserST, error) {
	entries, err := os.ReadDir(config.Get().Users.Path)
	if err != nil {
		return nil, err
	}
	users := make([]*UserST, 0, len(entries))
	var errs []error
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		ext := path.Ext(entry.Name())
		if ext != ".json" {
			continue
		}
		user, err := GetUserByUsername(strings.TrimSuffix(entry.Name(), ext))
		if err != nil {
			errs = append(errs, err)
			continue
		}
		if user == nil {
			continue
		}
		users = append(users, user)
	}
	return users, errors.Join(errs...)
}

func GetUserByUsername(username string) (*UserST, error) {
	var user UserST
	bytes, err := os.ReadFile(path.Join(config.Get().Users.Path, fmt.Sprintf("%s.json", username)))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUserPassword(username, password string) (*UserST, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	encryptedPassword, err := util.EncryptPassword(password)
	if err != nil {
		return nil, err
	}
	user.EncryptedPassword = encryptedPassword
	if err := writeUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

func writeUser(user *UserST) error {
	user.UpdatedAt = time.Now().UTC()
	bytes, err := json.MarshalIndent(user, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(path.Join(config.Get().Users.Path, fmt.Sprintf("%s.json", user.Username)), bytes, os.ModePerm)
}

func InitUsers() error {
	users, _ := GetUsers()
	if len(users) == 0 {
		encryptedPassword, err := util.EncryptPassword("password")
		if err != nil {
			return err
		}
		now := time.Now().UTC()
		adminUser := UserST{
			Username:          "admin",
			EncryptedPassword: encryptedPassword,
			CreatedAt:         now,
			UpdatedAt:         now,
		}
		if err := writeUser(&adminUser); err != nil {
			return err
		}
	}
	return nil
}
