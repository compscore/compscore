package round

import (
	"strconv"

	"github.com/compscore/compscore/pkg/data"
	"github.com/compscore/compscore/pkg/ent"
	"github.com/compscore/compscore/pkg/ent/round"
	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	roundStr := ctx.Param("round")

	round_number, err := strconv.Atoi(roundStr)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	entRound_i, err := data.Client(
		func(client *ent.Client) (interface{}, error) {
			return client.Round.Query().
				WithStatus(
					func(query *ent.StatusQuery) {
						query.WithTeam().WithCheck()
					},
				).
				Where(
					round.NumberEQ(round_number),
				).
				Only(ctx)
		},
	)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	entRound, ok := entRound_i.(*ent.Round)
	if !ok {
		ctx.JSON(500, gin.H{
			"error": "entRound_i.(*ent.Round) failed",
		})
		return
	}

	ctx.JSON(200, entRound)
}
