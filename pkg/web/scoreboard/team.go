package scoreboard

import (
	"fmt"
	"strconv"

	"github.com/compscore/compscore/pkg/cache"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Team(ctx *gin.Context) {
	teamStr := ctx.Param("team")

	if config.Production {
		cachedData, err := cache.Client.Get(ctx, fmt.Sprintf("scoreboard/team/%s", teamStr)).Result()
		if err != nil && err != redis.Nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if err == nil {
			ctx.JSON(200, cachedData)
			return
		}
	}

	team, err := strconv.Atoi(teamStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	teamScoreboard, err := data.Scoreboard.Team(team, 15)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if config.Production {
		err = cache.Client.Set(ctx, fmt.Sprintf("scoreboard/team/%s", teamStr), teamScoreboard, config.Redis.FastRefresh).Err()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	ctx.JSON(200, teamScoreboard)
}

func TeamRound(ctx *gin.Context) {
	teamStr := ctx.Param("team")
	roundStr := ctx.Param("round")

	if config.Production {
		cachedData, err := cache.Client.Get(ctx, fmt.Sprintf("scoreboard/team/%s/%s", teamStr, roundStr)).Result()
		if err != nil && err != redis.Nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if err == nil {
			ctx.JSON(200, cachedData)
			return
		}
	}

	team, err := strconv.Atoi(teamStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	round, err := strconv.Atoi(roundStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	teamScoreboard, err := data.Scoreboard.TeamRound(team, round, 15)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if config.Production {
		err = cache.Client.Set(ctx, fmt.Sprintf("scoreboard/team/%s/%s", teamStr, roundStr), teamScoreboard, config.Redis.FastRefresh).Err()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	ctx.JSON(200, teamScoreboard)
}
