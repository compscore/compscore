package web

import (
	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
)

func checkScoreboard(ctx *gin.Context) {
	check := ctx.Param("check")

	checkScoreboard, err := data.Status.CheckScoreboard(check, 10)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, checkScoreboard)
}
