package web

import (
	"strconv"

	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
)

func teamScoreboard(ctx *gin.Context) {
	teamStr := ctx.Param("team")
	team, err := strconv.Atoi(teamStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if 0 >= team && team >= 127 {
		ctx.JSON(400, gin.H{"error": "Team number be between 1 and 127"})
		return
	}

	teamScoreboard, err := data.Scoreboard.Team(int8(team), 10)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, teamScoreboard)
}
