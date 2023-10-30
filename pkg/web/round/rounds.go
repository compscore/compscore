package round

import (
	"github.com/compscore/compscore/pkg/cache"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Rounds(ctx *gin.Context) {
	if config.Production {
		cachedData, err := cache.Client.Get(ctx, "rounds").Result()
		if err != nil && err != redis.Nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err == nil {
			ctx.JSON(200, cachedData)
			return
		}
	}

	entRounds, err := data.Round.GetAll()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if config.Production {
		err = cache.Client.Set(ctx, "rounds", entRounds, config.Redis.FastRefresh).Err()
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	ctx.JSON(200, entRounds)
}
