package web

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/compscore/compscore/pkg/auth"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/data"
	"github.com/compscore/compscore/pkg/web/models"
	"github.com/gin-gonic/gin"
)

// login authenticates a user and returns a JWT
//
// @Summary Authenticate a user
// @Description Authenticate a user and return a JWT
// @Tags auth
// @Accept json
// @Produce json
// @Param body body models.Login true "Username and password"
// @Success 200 {object} models.Cookie
// @Failure 400 {object} models.Error
// @Router /api/login [post]
func login(ctx *gin.Context) {
	var body models.Login

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

	if body.Username == "" {
		ctx.JSON(
			http.StatusBadRequest,
			models.Error{
				Error: "no username provided",
			},
		)
		return
	}

	if body.Password == "" {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Error: "no password provided",
		})
		return
	}

	success, err := data.Team.CheckPasswordByName(body.Username, body.Password)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	if !success {
		ctx.JSON(
			http.StatusBadRequest,
			models.Error{
				Error: "invalid username or password",
			},
		)
		return
	}

	token, expiration, err := auth.GenerateJWT(body.Username)
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
