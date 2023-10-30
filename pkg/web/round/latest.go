package round

import (
	"encoding/json"

	"github.com/compscore/compscore/pkg/cache"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Latest(ctx *gin.Context) {
	if config.Production {
		cachedData, err := cache.Client.Get(ctx, "round/latest").Result()
		if err != nil && err != redis.Nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err == nil {
			ctx.String(200, cachedData)
			return
		}
	}

	round, err := data.Round.GetLastRound()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if config.Production {
		redisObject, err := json.Marshal(round)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = cache.Client.Set(ctx, "round/latest", string(redisObject), config.Redis.FastRefresh).Err()
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	ctx.JSON(200, round)
}
