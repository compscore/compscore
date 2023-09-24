package web

import (
	"fmt"
	"net/http"

	"github.com/compscore/compscore/pkg/auth"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
)

func login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	if username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "No username provided",
		})
		return
	}

	password := ctx.PostForm("password")
	if password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "No password provided",
		})
		return
	}

	success, err := data.Team.CheckPasswordByName(username, password)
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

	token, expiration, err := auth.GenerateJWT(username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.SetCookie("auth", token, expiration, "/", fmt.Sprintf("%s:%d", config.Web.Hostname, config.Web.Port), false, true)
	ctx.Status(http.StatusOK)
	// ctx.Redirect(302, "/")
}
