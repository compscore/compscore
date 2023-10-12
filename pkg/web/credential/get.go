package credential

import (
	"net/http"

	"github.com/compscore/compscore/pkg/auth"
	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	check := ctx.Param("check")
	entTeam, err := auth.Parse(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	entCredential, err := data.Credential.Get(entTeam.Number, check)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(
		http.StatusOK,
		entCredential,
	)
}
