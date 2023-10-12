package scoreboard

import (
	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
)

func Scoreboard(ctx *gin.Context) {
	scoreboard, err := data.Scoreboard.Main()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, scoreboard)
}
