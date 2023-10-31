package credential

import (
	"context"
	"net/http"

	"github.com/compscore/compscore/pkg/auth"
	"github.com/compscore/compscore/pkg/data"
	"github.com/compscore/compscore/pkg/ent"
	"github.com/compscore/compscore/pkg/ent/credential"
	"github.com/compscore/compscore/pkg/ent/team"
	"github.com/gin-gonic/gin"
)

func Credentials(ctx *gin.Context) {
	entTeam, err := auth.Parse(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	entCredentials_i, err := data.Client(
		func(client *ent.Client, ctx context.Context) (interface{}, error) {
			return client.Credential.Query().
				WithCheck().
				Where(
					credential.HasTeamWith(
						team.NumberEQ(entTeam.Number),
					),
				).
				All(ctx)
		},
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	entCredentials, ok := entCredentials_i.([]*ent.Credential)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to convert credentials",
		})
		return
	}

	ctx.JSON(http.StatusOK, entCredentials)
}
