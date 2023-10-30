package scoreboard

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

func Round(ctx *gin.Context) {
	roundStr := ctx.Param("round")

	if config.Production {
		cachedData, err := cache.Client.Get(ctx, fmt.Sprintf("scoreboard/round/%s", roundStr)).Result()
		if err != nil && err != redis.Nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if err == nil {
			ctx.String(200, cachedData)
			return
		}
	}

	round, err := strconv.Atoi(roundStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	scoreboard, err := data.Scoreboard.Round(round)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if config.Production {
		redisObject, err := json.Marshal(scoreboard)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		err = cache.Client.Set(ctx, fmt.Sprintf("scoreboard/round/%s", roundStr), string(redisObject), config.Redis.SlowRefresh).Err()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	ctx.JSON(200, scoreboard)
}
