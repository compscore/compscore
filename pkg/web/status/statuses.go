package status

import (
	"encoding/json"
	"net/http"

	"github.com/compscore/compscore/pkg/cache"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/data"
	"github.com/compscore/compscore/pkg/web/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// Statuses returns all status entries
//
// @Summary Get all status entries
// @Description Get all status entries
// @Tags status
// @Accept json
// @Produce json
// @Success 200 {array} models.Status
// @Failure 500 {object} models.Error
// @Router /statuses [get]
func Statuses(ctx *gin.Context) {
	if config.Production {
		cachedData, err := cache.Client.Get(ctx, "statuses").Result()
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

	entStatuses, err := data.Status.GetAllWithEdges()
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
		redisObject, err := json.Marshal(entStatuses)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				models.Error{
					Error: err.Error(),
				},
			)
			return
		}

		err = cache.Client.Set(ctx, "statuses", string(redisObject), config.Redis.SlowRefresh).Err()
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

	ctx.JSON(http.StatusOK, entStatuses)
}
