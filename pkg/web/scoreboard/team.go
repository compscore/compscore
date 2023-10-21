package scoreboard

import (
	"strconv"

	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
)

func Team(ctx *gin.Context) {
	teamStr := ctx.Param("team")
	team, err := strconv.Atoi(teamStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	teamScoreboard, err := data.Scoreboard.Team(team, 10)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, teamScoreboard)
}

func TeamRound(ctx *gin.Context) {
	teamStr := ctx.Param("team")
	roundStr := ctx.Param("round")

	team, err := strconv.Atoi(teamStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	round, err := strconv.Atoi(roundStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	teamScoreboard, err := data.Scoreboard.TeamRound(team, round, 10)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, teamScoreboard)
}
