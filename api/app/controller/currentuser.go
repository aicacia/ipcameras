package controller

import (
	"log"
	"net/http"
	"strings"

	"github.com/aicacia/ipcameras/api/app/middleware"
	"github.com/aicacia/ipcameras/api/app/model"
	"github.com/aicacia/ipcameras/api/app/repo"
	"github.com/gofiber/fiber/v2"
)

// GetCurrentUser
//
//	@ID				  current-user
//	@Summary		Get current user
//	@Tags			  current-user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}  model.UserST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user [get]
//
//	@Security		Authorization
func GetCurrentUser(c *fiber.Ctx) error {
	user := middleware.GetUser(c)
	return c.JSON(model.UserFromRepo(user))
}

// PatchResetPassword
//
//	@Summary		Resets a user's password
//	@ID				  reset-password
//	@Tags		  	current-user
//	@Accept			json
//	@Produce		json
//	@Param			resetPassword	body    model.ResetPasswordST	true	"reset user's password"
//	@Success		204
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user/reset-password [patch]
//
//	@Security		Authorization
func PatchResetPassword(c *fiber.Ctx) error {
	var resetPassword model.ResetPasswordST
	if err := c.BodyParser(&resetPassword); err != nil {
		log.Printf("failed to parse reset password: %v\n", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid")
	}
	password := strings.TrimSpace(resetPassword.Password)
	passwordConfirmation := strings.TrimSpace(resetPassword.PasswordConfirmation)
	errors := model.NewError(http.StatusBadRequest)
	if password != passwordConfirmation {
		errors.AddError("password_confirmation", "mismatch", "body")
	}
	if errors.HasErrors() {
		return errors
	}
	user := middleware.GetUser(c)
	_, err := repo.UpdateUserPassword(user.Username, password)
	if err != nil {
		log.Printf("failed to update user password: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	c.Status(http.StatusNoContent)
	return c.Send(nil)
}
