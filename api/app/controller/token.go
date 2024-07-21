package controller

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/aicacia/ipcameras/api/app/config"
	"github.com/aicacia/ipcameras/api/app/jwt"
	"github.com/aicacia/ipcameras/api/app/model"
	"github.com/aicacia/ipcameras/api/app/service"
	"github.com/aicacia/ipcameras/api/app/util"
	"github.com/gofiber/fiber/v2"
)

// PostToken
//
//	@ID				  token
//	@Summary		create a token by authenticating a user
//	@Tags			  token
//	@Accept			json
//	@Produce		json
//	@Param      credentials body     model.CredentialsST true "user credentials"
//	@Success		201	{object}	model.TokenST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/token [post]
func PostToken(c *fiber.Ctx) error {
	var credentials model.CredentialsST
	if err := c.BodyParser(&credentials); err != nil {
		slog.Error("error parsing credentials", "error", err)
		return model.NewError(http.StatusBadRequest).AddError("credentials", "invalid")
	}
	user, err := service.GetUserByUsername(credentials.Username)
	if err != nil {
		slog.Error("error getting user by username", "error", err)
		return model.NewError(http.StatusBadRequest).AddError("credentials", "invalid")
	}
	if user == nil {
		slog.Error("user not found", "username", credentials.Username)
		return model.NewError(http.StatusBadRequest).AddError("credentials", "invalid")
	}
	if ok, err := util.VerifyPassword(credentials.Password, user.EncryptedPassword); !ok || err != nil {
		slog.Error("invalid password", "error", err)
		return model.NewError(http.StatusBadRequest).AddError("credentials", "invalid")
	}
	return sendToken(c, "password", user)
}

func sendToken(
	c *fiber.Ctx,
	issuedTokenType string,
	user *service.UserST,
) error {
	now := time.Now().UTC()
	claims := jwt.Claims{
		Type:             jwt.BearerTokenType,
		Subject:          user.Username,
		NotBeforeSeconds: now.Unix(),
		IssuedAtSeconds:  now.Unix(),
		ExpiresAtSeconds: now.Unix() + int64(config.Get().JWT.ExpiresInSeconds),
		Issuer:           config.Get().URL,
	}
	accessToken, err := jwt.CreateToken(&claims, config.Get().JWT.Secret)
	if err != nil {
		slog.Error("failed to create access token", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	refreshToken, err := jwt.CreateToken(claims.ToRefreshClaims(now.Unix()+int64(config.Get().JWT.RefreshExpiresInSeconds)), config.Get().JWT.Secret)
	if err != nil {
		slog.Error("failed to create refresh token", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	return c.JSON(model.TokenST{
		AccessToken:           accessToken,
		TokenType:             "Bearer",
		IssuedTokenType:       issuedTokenType,
		ExpiresIn:             config.Get().JWT.ExpiresInSeconds,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresIn: config.Get().JWT.RefreshExpiresInSeconds,
	})
}
