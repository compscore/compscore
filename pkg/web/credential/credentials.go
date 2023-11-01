package credential

import (
	"context"
	"net/http"

	"github.com/compscore/compscore/pkg/auth"
	"github.com/compscore/compscore/pkg/data"
	"github.com/compscore/compscore/pkg/ent"
	"github.com/compscore/compscore/pkg/ent/credential"
	"github.com/compscore/compscore/pkg/ent/team"
	"github.com/compscore/compscore/pkg/web/models"
	"github.com/gin-gonic/gin"
)

// Credentials returns all credentials for a team
//
// @Summary Get all credentials for a team
// @Description Get all credentials for a team
// @Tags credential
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} []models.Credential
// @Failure 401 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/credentials [get]
func Credentials(ctx *gin.Context) {
	entTeam, err := auth.Parse(ctx)
	if err != nil {
		ctx.JSON(
			http.StatusUnauthorized,
			models.Error{
				Error: err.Error(),
			},
		)
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
		ctx.JSON(
			http.StatusInternalServerError,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	entCredentials, ok := entCredentials_i.([]*ent.Credential)
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			models.Error{
				Error: "failed to assert credentials type",
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, entCredentials)
}
