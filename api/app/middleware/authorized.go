package middleware

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/aicacia/ipcameras/api/app/config"
	"github.com/aicacia/ipcameras/api/app/jwt"
	"github.com/aicacia/ipcameras/api/app/model"
	"github.com/aicacia/ipcameras/api/app/service"
	"github.com/gofiber/fiber/v2"
)

var claimsLocalKey = "claims"
var userLocalKey = "user"

func AuthorizedMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, tokenString := GetAuthorizationFromContext(c)
		claims, err := jwt.ParseClaimsFromToken(tokenString, config.Get().JWT.Secret)
		if err != nil {
			slog.Error("failed to parse claims from token", "error", err)
			return model.NewError(http.StatusUnauthorized).AddError("authorization", "invalid")
		}
		if claims.Type != jwt.BearerTokenType {
			return model.NewError(http.StatusUnauthorized).AddError("authorization", "invalid")
		}
		user, err := service.GetUserByUsername(claims.Subject)
		if err != nil {
			slog.Error("failed to fetch user", "error", err)
			return model.NewError(http.StatusUnauthorized).AddError("authorization", "invalid")
		}
		c.Locals(userLocalKey, user)
		c.Locals(claimsLocalKey, claims)
		return c.Next()
	}
}

func GetClaims(c *fiber.Ctx) *jwt.Claims {
	claims := c.Locals(claimsLocalKey)
	return claims.(*jwt.Claims)
}

func GetUser(c *fiber.Ctx) *service.UserST {
	user := c.Locals(userLocalKey)
	return user.(*service.UserST)
}

func GetAuthorizationFromContext(c *fiber.Ctx) (string, string) {
	authorizationHeader := strings.TrimSpace(c.Get("Authorization"))
	if len(authorizationHeader) != 0 {
		parts := strings.SplitN(authorizationHeader, " ", 2)
		if len(parts) == 2 {
			tokenType := strings.TrimSpace(parts[0])
			token := strings.TrimSpace(parts[1])
			return tokenType, token
		}
	}
	return "", ""
}
