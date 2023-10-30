package check

import (
	"encoding/json"

	"github.com/compscore/compscore/pkg/cache"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Checks(ctx *gin.Context) {
	if config.Production {
		cachedData, err := cache.Client.Get(ctx, "checks").Result()
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

	entChecks, err := data.Check.GetAll()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if config.Production {
		redisObject, err := json.Marshal(entChecks)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = cache.Client.Set(ctx, "checks", string(redisObject), config.Redis.SlowRefresh).Err()
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	ctx.JSON(200, entChecks)
}
