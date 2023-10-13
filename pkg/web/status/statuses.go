package status

import (
	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
)

func Statuses(ctx *gin.Context) {
	entStatuses, err := data.Status.GetAllWithEdges()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, entStatuses)
}
