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

func GetByTeamCheck(ctx *gin.Context) {
	teamStr := ctx.Param("team")
	check := ctx.Param("check")

	if config.Production {
		cachedData, err := cache.Client.Get(ctx, fmt.Sprintf("status/team/%s/check/%s", check, teamStr)).Result()
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

	entStatus, err := data.Status.GetAllByCheckAndTeamWithEdges(check, team)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if config.Production {
		err = cache.Client.Set(ctx, fmt.Sprintf("status/team/%s/check/%s", check, teamStr), entStatus, config.Redis.FastRefresh).Err()
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	ctx.JSON(200, entStatus)
}
