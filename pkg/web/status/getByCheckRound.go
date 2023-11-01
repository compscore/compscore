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

// GetByCheckRound returns all status entries for a given check and round
//
// @Summary Get all status entries for a given check and round
// @Description Get all status entries for a given check and round
// @Tags status
// @Accept json
// @Produce json
// @Param check path string true "Check name"
// @Param round path int true "Round number"
// @Success 200 {array} models.Status
// @Failure 400 {object} models.Error
// @Router /status/check/{check}/round/{round} [get]
func GetByCheckRound(ctx *gin.Context) {
	check := ctx.Param("check")
	roundStr := ctx.Param("round")

	if config.Production {
		cachedData, err := cache.Client.Get(ctx, fmt.Sprintf("status/check/%s/round/%s", check, roundStr)).Result()
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

	round_number, err := strconv.Atoi(roundStr)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	entStatus, err := data.Status.GetAllByRoundAndCheckWithEdges(round_number, check)
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

		err = cache.Client.Set(ctx, fmt.Sprintf("status/check/%s/round/%s", check, roundStr), string(redisObject), config.Redis.FastRefresh).Err()
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
