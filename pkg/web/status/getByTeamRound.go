package status

import (
	"strconv"

	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
)

func GetByTeamRound(ctx *gin.Context) {
	teamStr := ctx.Param("team")
	roundStr := ctx.Param("round")

	team, err := strconv.Atoi(teamStr)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	round, err := strconv.Atoi(roundStr)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	entStatus, err := data.Status.GetAllByRoundAndTeamWithEdges(round, team)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, entStatus)
}
