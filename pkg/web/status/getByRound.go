package status

import (
	"encoding/json"
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
		redisObject, err := json.Marshal(entStatus)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = cache.Client.Set(ctx, fmt.Sprintf("status/round/%s", roundStr), string(redisObject), config.Redis.FastRefresh).Err()
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	ctx.JSON(200, entStatus)
}
