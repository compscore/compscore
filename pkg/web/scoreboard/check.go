package scoreboard

import (
	"strconv"

	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
)

func Check(ctx *gin.Context) {
	check := ctx.Param("check")

	checkScoreboard, err := data.Scoreboard.Check(check, 10)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, checkScoreboard)
}

func CheckRound(ctx *gin.Context) {
	check := ctx.Param("check")
	roundStr := ctx.Param("round")

	round, err := strconv.Atoi(roundStr)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	checkScoreboard, err := data.Scoreboard.CheckRound(check, round, 10)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, checkScoreboard)
}
