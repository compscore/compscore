package web

import (
	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
)

func login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	if username == "" {
		ctx.JSON(200, gin.H{
			"error": "No username provided",
		})
		return
	}

	password := ctx.PostForm("password")
	if password == "" {
		ctx.JSON(200, gin.H{
			"error": "No password provided",
		})
		return
	}

	success, err := data.Team.CheckPasswordByName(username, password)
	if err != nil {
		ctx.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !success {
		ctx.JSON(200, gin.H{
			"error": "Invalid username or password",
		})
		return
	} else {
		ctx.SetCookie("username", username, 0, "/", "", false, true)
		ctx.Redirect(302, "/")
	}
}
