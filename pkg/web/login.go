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

type login_s struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(ctx *gin.Context) {
	var body login_s

	body_bytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	json.Unmarshal(body_bytes, &body)

	if body.Username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "No username provided",
		})
		return
	}

	if body.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "No password provided",
		})
		return
	}

	success, err := data.Team.CheckPasswordByName(body.Username, body.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !success {
		ctx.JSON(400, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	token, expiration, err := auth.GenerateJWT(body.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"name":       "auth",
			"token":      token,
			"expiration": expiration,
			"path":       "/",
			"domain":     config.Web.Hostname,
			"secure":     false,
			"httponly":   true,
		},
	)
}
