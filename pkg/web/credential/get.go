package credential

import (
	"net/http"

	"github.com/compscore/compscore/pkg/auth"
	"github.com/compscore/compscore/pkg/data"
	"github.com/compscore/compscore/pkg/web/models"
	"github.com/gin-gonic/gin"
)

// Get returns a credential
//
// @Summary Get a credential
// @Description Get a credential
// @Tags credential
// @Accept json
// @Produce json
// @Security ServiceAuth
// @Param check path string true "Check name"
// @Success 200 {object} models.Credential
// @Failure 400 {object} models.Error
// @Failure 401 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /credential/{check} [get]
func Get(ctx *gin.Context) {
	check := ctx.Param("check")
	entTeam, err := auth.Parse(ctx)
	if err != nil {
		ctx.JSON(
			http.StatusUnauthorized,
			models.Error{
				Error: "user not authenticated",
			},
		)
		return
	}

	entCredential, err := data.Credential.Get(entTeam.Number, check)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		entCredential,
	)
}
