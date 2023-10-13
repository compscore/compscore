package check

import (
	"github.com/compscore/compscore/pkg/auth"
	"github.com/compscore/compscore/pkg/data"
	"github.com/compscore/compscore/pkg/ent"
	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	check := ctx.Param("check")

	var entCheck *ent.Check
	var err error

	entTeam, err := auth.Parse(ctx)
	if err != nil {
		entCheck, err = data.Check.Get(check)
	} else {
		entCheck, err = data.Check.GetWithTeamCredenital(check, entTeam.Number)
	}

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, entCheck)

}
