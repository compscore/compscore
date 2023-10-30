package team

import (
	"fmt"
	"strconv"

	"github.com/compscore/compscore/pkg/cache"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Get(ctx *gin.Context) {
	teamStr := ctx.Param("team")

	if config.Production {
		cachedData, err := cache.Client.Get(ctx, fmt.Sprintf("team/%s", teamStr)).Result()
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

	team, err := strconv.Atoi(teamStr)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	entTeam, err := data.Team.Get(team)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if config.Production {
		err = cache.Client.Set(ctx, fmt.Sprintf("team/%s", teamStr), entTeam, config.Redis.SlowRefresh).Err()
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	ctx.JSON(200, entTeam)
}
