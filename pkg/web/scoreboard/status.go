package scoreboard

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/compscore/compscore/pkg/cache"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/data"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Status(ctx *gin.Context) {
	check := ctx.Param("check")
	teamStr := ctx.Param("team")

	if config.Production {
		cachedData, err := cache.Client.Get(ctx, fmt.Sprintf("scoreboard/status/%s/%s", check, teamStr)).Result()
		if err != nil && err != redis.Nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err == nil {
			ctx.JSON(http.StatusOK, cachedData)
			return
		}
	}

	team, err := strconv.Atoi(teamStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid team"})
		return
	}

	statusHistory, err := data.Scoreboard.History(check, team, 10)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if config.Production {
		err = cache.Client.Set(ctx, fmt.Sprintf("scoreboard/status/%s/%s", check, teamStr), statusHistory, config.Redis.FastRefresh).Err()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusOK, statusHistory)
}

func StatusRound(ctx *gin.Context) {
	check := ctx.Param("check")
	teamStr := ctx.Param("team")
	roundStr := ctx.Param("round")

	if config.Production {
		cachedData, err := cache.Client.Get(ctx, fmt.Sprintf("scoreboard/status/%s/%s/%s", check, teamStr, roundStr)).Result()
		if err != nil && err != redis.Nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err == nil {
			ctx.JSON(http.StatusOK, cachedData)
			return
		}
	}

	team, err := strconv.Atoi(teamStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid team"})
		return
	}

	round, err := strconv.Atoi(roundStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid round"})
		return
	}

	statusHistory, err := data.Scoreboard.HistoryRound(check, team, round, 10)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if config.Production {
		err = cache.Client.Set(ctx, fmt.Sprintf("scoreboard/status/%s/%s/%s", check, teamStr, roundStr), statusHistory, config.Redis.FastRefresh).Err()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusOK, statusHistory)
}
