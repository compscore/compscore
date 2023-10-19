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

	statusHistory, err := data.Scoreboard.History(check, team, 10)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, statusHistory)
}

func StatusRound(ctx *gin.Context) {
	check := ctx.Param("check")
	teamStr := ctx.Param("team")
	roundStr := ctx.Param("round")

	team, err := strconv.Atoi(teamStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid team"})
		return
	}

	round, err := strconv.Atoi(roundStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid round"})
		return
	}

	statusHistory, err := data.Scoreboard.HistoryRound(check, team, round, 10)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, statusHistory)
}
