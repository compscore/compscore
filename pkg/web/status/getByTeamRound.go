package status

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/compscore/compscore/pkg/cache"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/data"
	"github.com/compscore/compscore/pkg/web/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// GetByTeamRound returns all status entries for a given team and round
//
// @Summary Get all status entries for a given team and round
// @Description Get all status entries for a given team and round
// @Tags status
// @Accept json
// @Produce json
// @Param team path int true "Team ID"
// @Param round path int true "Round number"
// @Success 200 {array} models.Status
// @Failure 500 {object} models.Error
// @Router /status/team/{team}/round/{round} [get]
func GetByTeamRound(ctx *gin.Context) {
	teamStr := ctx.Param("team")
	roundStr := ctx.Param("round")

	if config.Production {
		cachedData, err := cache.Client.Get(ctx, fmt.Sprintf("status/team/%s/round/%s", teamStr, roundStr)).Result()
		if err != nil && err != redis.Nil {
			ctx.JSON(
				http.StatusInternalServerError,
				models.Error{
					Error: err.Error(),
				},
			)
			return
		}

		if err == nil {
			ctx.String(http.StatusOK, cachedData)
			return
		}
	}

	team, err := strconv.Atoi(teamStr)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	round, err := strconv.Atoi(roundStr)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	entStatus, err := data.Status.GetAllByRoundAndTeamWithEdges(round, team)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	if config.Production {
		redisObject, err := json.Marshal(entStatus)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				models.Error{
					Error: err.Error(),
				},
			)
			return
		}

		err = cache.Client.Set(ctx, fmt.Sprintf("status/team/%s/round/%s", teamStr, roundStr), string(redisObject), config.Redis.FastRefresh).Err()
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				models.Error{
					Error: err.Error(),
				},
			)
			return
		}
	}

	ctx.JSON(http.StatusOK, entStatus)
}
