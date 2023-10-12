package round

import (
	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
)

func Rounds(ctx *gin.Context) {
	entRounds, err := data.Round.GetAll()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, entRounds)
}
