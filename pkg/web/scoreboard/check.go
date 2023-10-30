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

func Check(ctx *gin.Context) {
	check := ctx.Param("check")

	if config.Production {
		cachedData, err := cache.Client.Get(ctx, fmt.Sprintf("scoreboard/check/%s", check)).Result()
		if err != nil && err != redis.Nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if err == nil {
			ctx.String(200, cachedData)
			return
		}
	}

	checkScoreboard, err := data.Scoreboard.Check(check, 10)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if config.Production {
		redisObject, err := json.Marshal(checkScoreboard)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		err = cache.Client.Set(ctx, fmt.Sprintf("scoreboard/check/%s", check), string(redisObject), config.Redis.FastRefresh).Err()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	ctx.JSON(200, checkScoreboard)
}

func CheckRound(ctx *gin.Context) {
	check := ctx.Param("check")
	roundStr := ctx.Param("round")

	if config.Production {
		cachedData, err := cache.Client.Get(ctx, fmt.Sprintf("scoreboard/check/%s/%s", check, roundStr)).Result()
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
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	checkScoreboard, err := data.Scoreboard.CheckRound(check, round, 10)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if config.Production {
		redisObject, err := json.Marshal(checkScoreboard)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		err = cache.Client.Set(ctx, fmt.Sprintf("scoreboard/check/%s/%s", check, roundStr), string(redisObject), config.Redis.FastRefresh).Err()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	ctx.JSON(200, checkScoreboard)
}
