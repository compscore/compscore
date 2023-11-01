package web

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/compscore/compscore/pkg/auth"
	"github.com/compscore/compscore/pkg/data"
	"github.com/compscore/compscore/pkg/web/models"
	"github.com/gin-gonic/gin"
)

// password changes a user's password
// @Summary Change a user's password
// @Description Change a user's password
// @Tags auth
// @Accept json
// @Produce json
// @Param body body models.ChangePassword true "Old and new password"
// @Success 200 "Password changed"
// @Failure 400 {object} models.Error
// @Failure 401 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /password [post]
func password(ctx *gin.Context) {
	entTeam, err := auth.Parse(ctx)
	if err != nil {
		ctx.JSON(
			http.StatusUnauthorized,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	var body models.ChangePassword

	body_bytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	err = json.Unmarshal(body_bytes, &body)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	correctPassword, err := data.Team.CheckPassword(entTeam.Number, body.OldPassword)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	if !correctPassword {
		ctx.JSON(
			http.StatusBadRequest,
			models.Error{
				Error: "incorrect password",
			},
		)
		return
	}

	_, err = data.Team.UpdatePassword(entTeam, body.NewPassword)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	ctx.Status(http.StatusOK)
}
