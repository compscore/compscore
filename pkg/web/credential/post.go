package credential

import (
	"io"
	"net/http"

	"encoding/json"

	"github.com/compscore/compscore/pkg/auth"
	"github.com/compscore/compscore/pkg/data"
	"github.com/compscore/compscore/pkg/web/models"
	"github.com/gin-gonic/gin"
)

// Post updates a credential
//
// @Summary Update a credential
// @Description Update a credential
// @Tags credential
// @Accept json
// @Produce json
// @Security ServiceAuth
// @Param check path string true "Check name"
// @Param body body models.CredentialEdit true "New password"
// @Success 200
// @Failure 400 {object} models.Error
// @Failure 401 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /credential/{check} [post]
func Post(ctx *gin.Context) {
	var body models.CredentialEdit

	check := ctx.Param("check")
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

	body_bytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	err = json.Unmarshal(body_bytes, &body)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	_, err = data.Credential.UpdatePassword(entTeam.Number, check, body.Password)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	ctx.Status(http.StatusOK)
}
