package status

import (
	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
)

func GetByCheck(ctx *gin.Context) {
	check := ctx.Param("check")

	entStatus, err := data.Status.GetAllByCheckWithEdges(check)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, entStatus)
}
