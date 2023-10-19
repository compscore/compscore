package status

import (
	"strconv"

	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	teamStr := ctx.Param("team")
	check := ctx.Param("check")
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

	entStatus, err := data.Status.GetWithEdges(team, check, round)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, entStatus)
}
