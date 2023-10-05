package web

import (
	"strconv"

	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
)

func roundScoreboard(ctx *gin.Context) {
	roundStr := ctx.Param("round")
	round, err := strconv.Atoi(roundStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	scoreboard, err := data.Scoreboard.Round(round)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, scoreboard)
}
