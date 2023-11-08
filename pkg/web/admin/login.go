package admin

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/compscore/compscore/pkg/auth"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/ent/team"
	"github.com/compscore/compscore/pkg/web/models"
	"github.com/gin-gonic/gin"
)

// Login authenticates into another team from admin account and returns a JWT
//
// @Summary Authenticate into another team
// @Description Authenticate into another team and return a JWT
// @Tags admin
// @Accept json
// @Produce json
// @Security ServiceAuth
// @Param body body models.AdminLogin true "Team name"
// @Success 200 {object} models.Cookie
// @Failure 400 {object} models.Error
// @Router /admin/login [post]
func Login(ctx *gin.Context) {
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

	var body models.AdminLogin

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

	token, expiration, err := auth.GenerateJWT(body.Team)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		models.Cookie{
			Name:       "auth",
			Token:      token,
			Expiration: expiration,
			Path:       "/",
			Domain:     config.Web.Hostname,
			Secure:     false,
			HttpOnly:   true,
		},
	)
}
