package web

import (
	"net/http"

	"github.com/compscore/compscore/pkg/auth"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/data"
	"github.com/compscore/compscore/pkg/structs"
	"github.com/gin-gonic/gin"
)

func checks(ctx *gin.Context) {
	entTeam, err := auth.Parse(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	credentials := make([]structs.Credential, len(config.Checks))

	for i, check := range config.Checks {
		entCredential, err := data.Credential.Get(entTeam.Number, check.Name)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		credentials[i] = structs.Credential{
			Check:    check.Name,
			Password: entCredential.Password,
		}
	}

	ctx.JSON(http.StatusOK, credentials)
}
