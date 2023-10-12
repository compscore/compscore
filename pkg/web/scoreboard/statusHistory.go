package scoreboard

import (
	"net/http"
	"strconv"

	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
)

func Status(ctx *gin.Context) {
	check := ctx.Param("check")
	teamStr := ctx.Param("team")

	team, err := strconv.Atoi(teamStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid team"})
		return
	}

	if 0 >= team && team >= 127 {
		ctx.JSON(400, gin.H{"error": "Team number be between 1 and 127"})
		return
	}

	statusHistory, err := data.Scoreboard.History(check, int8(team), 10)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, statusHistory)
}
