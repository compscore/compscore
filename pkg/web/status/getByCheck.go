package status

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/compscore/compscore/pkg/cache"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/data"
	"github.com/compscore/compscore/pkg/web/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// GetByCheck returns all status entries for a given check
//
// @Summary Get all status entries for a given check
// @Description Get all status entries for a given check
// @Tags status
// @Accept json
// @Produce json
// @Param check path string true "Check name"
// @Success 200 {array} models.Status
// @Failure 400 {object} models.Error
// @Router /status/check/{check} [get]
func GetByCheck(ctx *gin.Context) {
	check := ctx.Param("check")

	if config.Production {
		cachedData, err := cache.Client.Get(ctx, fmt.Sprintf("status/check/%s", check)).Result()
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

	entStatus, err := data.Status.GetAllByCheckWithEdges(check)
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

		err = cache.Client.Set(ctx, fmt.Sprintf("status/check/%s", check), string(redisObject), config.Redis.FastRefresh).Err()
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
