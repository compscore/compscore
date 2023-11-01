package web

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/compscore/compscore/pkg/auth"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
)

// login_s is the struct used to unmarshal the JSON body of the login request
// @Summary body of login request
// @Description body of login request
// @Tags auth
type login_s struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// cookie_s is the struct used to marshal the JSON response of the login request
// @Summary response of login request
// @Description response of login request
// @Tags auth
type cookie_s struct {
	Name       string `json:"name"`
	Token      string `json:"token"`
	Expiration int    `json:"expiration"`
	Path       string `json:"path"`
	Domain     string `json:"domain"`
	Secure     bool   `json:"secure"`
	HttpOnly   bool   `json:"httponly"`
}

type error_s struct {
	Error string `json:"error"`
}

// login authenticates a user and returns a JWT
//
// @Summary Authenticate a user
// @Description Authenticate a user and return a JWT
// @Tags auth
// @Accept json
// @Produce json
// @Param body body login_s true "Username and password"
// @Success 200 {object} cookie_s
// @Failure 400
// @Header 200 {string} Set-Cookie "auth=JWT; Path=/; Domain=hostname; Secure; HttpOnly"
// @Router /api/login [post]
func login(ctx *gin.Context) {
	var body login_s

	body_bytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			error_s{
				Error: err.Error(),
			},
		)
		return
	}

	err = json.Unmarshal(body_bytes, &body)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			error_s{
				Error: err.Error(),
			},
		)
		return
	}

	if body.Username == "" {
		ctx.JSON(
			http.StatusBadRequest,
			error_s{
				Error: "no username provided",
			},
		)
		return
	}

	if body.Password == "" {
		ctx.JSON(http.StatusBadRequest, error_s{
			Error: "no password provided",
		})
		return
	}

	success, err := data.Team.CheckPasswordByName(body.Username, body.Password)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			error_s{
				Error: err.Error(),
			},
		)
		return
	}

	if !success {
		ctx.JSON(
			http.StatusBadRequest,
			error_s{
				Error: "invalid username or password",
			},
		)
		return
	}

	token, expiration, err := auth.GenerateJWT(body.Username)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			error_s{
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		cookie_s{
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
