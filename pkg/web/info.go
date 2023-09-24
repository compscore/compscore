package web

import (
	"net/http"

	"github.com/compscore/compscore/pkg/auth"
	"github.com/gin-gonic/gin"
)

func info(ctx *gin.Context) {
	entTeam, err := auth.Parse(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"team":   entTeam.Name,
		"number": entTeam.Number,
	})
}
