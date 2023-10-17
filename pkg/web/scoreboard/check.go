package scoreboard

import (
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
