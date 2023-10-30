package status

import (
	"fmt"
	"strconv"

	"github.com/compscore/compscore/pkg/cache"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func GetByRound(ctx *gin.Context) {
	roundStr := ctx.Param("round")

	if config.Production {
		cachedData, err := cache.Client.Get(ctx, fmt.Sprintf("status/round/%s", roundStr)).Result()
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

	round, err := strconv.Atoi(roundStr)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	entStatus, err := data.Status.GetAllByRoundWithEdges(round)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if config.Production {
		err = cache.Client.Set(ctx, fmt.Sprintf("status/round/%s", roundStr), entStatus, config.Redis.FastRefresh).Err()
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	ctx.JSON(200, entStatus)
}
