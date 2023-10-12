package check

import (
	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	check := ctx.Param("check")

	entCheck, err := data.Check.Get(check)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, entCheck)
}
