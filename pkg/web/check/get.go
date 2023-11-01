package check

import (
	"net/http"

	"github.com/compscore/compscore/pkg/auth"
	"github.com/compscore/compscore/pkg/data"
	"github.com/compscore/compscore/pkg/ent"
	"github.com/compscore/compscore/pkg/web/models"
	"github.com/gin-gonic/gin"
)

// Get returns a check
//
// @Summary Get a check
// @Description Get a check
// @Tags check
// @Accept json
// @Produce json
// @Param check path string true "Check ID"
// @Success 200 {object} models.Check
// @Failure 500 {object} models.Error
// @Router /check/{check} [get]
func Get(ctx *gin.Context) {
	check := ctx.Param("check")

	var entCheck *ent.Check
	var err error

	entTeam, err := auth.Parse(ctx)
	if err != nil {
		entCheck, err = data.Check.Get(check)
	} else {
		entCheck, err = data.Check.GetWithTeamCredential(check, entTeam.Number)
	}
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(200, entCheck)

}
