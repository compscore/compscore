package status

import (
	"strconv"

	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
)

func GetByTeamCheck(ctx *gin.Context) {
	teamStr := ctx.Param("team")
	check := ctx.Param("check")

	team, err := strconv.Atoi(teamStr)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	entStatus, err := data.Status.GetAllByCheckAndTeamWithEdges(check, team)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, entStatus)
}
