package scoreboard

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

func Scoreboard(ctx *gin.Context) {
	if config.Production {
		cachedData, err := cache.Client.Get(ctx, "scoreboard").Result()
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
			ctx.String(200, cachedData)
			return
		}
	}

	scoreboard, err := data.Scoreboard.Main()
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
		redisObject, err := json.Marshal(scoreboard)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				models.Error{
					Error: err.Error(),
				},
			)
			return
		}

		err = cache.Client.Set(ctx, "scoreboard", string(redisObject), config.Redis.MediumRefresh).Err()
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

	ctx.JSON(200, scoreboard)
}
