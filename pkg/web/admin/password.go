package admin

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/compscore/compscore/pkg/auth"
	"github.com/compscore/compscore/pkg/data"
	"github.com/compscore/compscore/pkg/ent/team"
	"github.com/compscore/compscore/pkg/web/models"
	"github.com/gin-gonic/gin"
)

// Password resets the password of another team from admin account
//
// @Summary Reset password of another team
// @Description Reset password of another team
// @Tags admin
// @Accept json
// @Produce json
// @Param body body models.AdminPasswordReset true "Team name and new password"
// @Success 200
// @Failure 400 {object} models.Error
// @Failure 401 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /admin/password [post]
func Password(ctx *gin.Context) {
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

	if entTeam.Role != team.RoleAdmin {
		ctx.JSON(
			http.StatusUnauthorized,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	var body models.AdminPasswordReset

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

	entTargetTeam, err := data.Team.GetByName(body.Team)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	_, err = data.Team.UpdatePassword(entTargetTeam, body.Password)
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
