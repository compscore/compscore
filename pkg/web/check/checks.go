package check

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

// Checks returns all checks
//
// @Summary Get all checks
// @Description Get all checks
// @Tags check
// @Accept json
// @Produce json
// @Success 200 {object} []models.Check
// @Failure 500 {object} models.Error
// @Router /checks [get]
func Checks(ctx *gin.Context) {
	if config.Production {
		cachedData, err := cache.Client.Get(ctx, "checks").Result()
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

	entChecks, err := data.Check.GetAll()
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
		redisObject, err := json.Marshal(entChecks)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				models.Error{
					Error: err.Error(),
				},
			)
			return
		}

		err = cache.Client.Set(ctx, "checks", string(redisObject), config.Redis.SlowRefresh).Err()
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

	ctx.JSON(http.StatusOK, entChecks)
}
