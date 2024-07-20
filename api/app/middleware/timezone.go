package middleware

import (
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/aicacia/ipcameras/api/app/model"
	"github.com/gofiber/fiber/v2"
)

var timezoneKey = "timezone"

func TimezoneMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		timezone, err := GetTimezoneFromContext(c)
		if err != nil {
			slog.Error("failed to get timezone", "error", err)
			return model.NewError(http.StatusBadRequest).AddError("x-timezone", "invalid")
		}
		c.Locals(timezoneKey, timezone)
		return c.Next()
	}
}

func GetTimezone(c *fiber.Ctx) *time.Location {
	return c.Locals(timezoneKey).(*time.Location)
}

func GetTimezoneFromContext(c *fiber.Ctx) (*time.Location, error) {
	timezoneString := strings.TrimSpace(c.Get("X-Timezone"))
	timezone, err := time.LoadLocation(timezoneString)
	if err != nil {
		return nil, err
	}
	return timezone, nil
}
