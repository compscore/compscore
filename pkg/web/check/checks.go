package check

import (
	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
)

func Checks(ctx *gin.Context) {
	entChecks, err := data.Check.GetAll()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, entChecks)
}
