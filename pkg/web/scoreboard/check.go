package scoreboard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/compscore/compscore/pkg/cache"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/data"
	"github.com/compscore/compscore/pkg/web/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// Check returns the scoreboard for a given check
//
// @Summary Get the scoreboard for a given check
// @Description Get the scoreboard for a given check
// @Tags scoreboard
// @Accept json
// @Produce json
// @Param check path string true "Check name"
// @Success 200 {object} structs.CheckScoreboard
// @Failure 500 {object} models.Error
// @Router /scoreboard/check/{check} [get]
func Check(ctx *gin.Context) {
	check := ctx.Param("check")

	if config.Production {
		cachedData, err := cache.Client.Get(ctx, fmt.Sprintf("scoreboard/check/%s", check)).Result()
		if err != nil && err != redis.Nil {
			ctx.JSON(
				http.StatusInternalServerError,
				models.Error{
					Error: err.Error(),
				},
			)
			return
		}

		if err == nil {
			ctx.String(http.StatusOK, cachedData)
			return
		}
	}

	checkScoreboard, err := data.Scoreboard.Check(check, 10)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	if config.Production {
		redisObject, err := json.Marshal(checkScoreboard)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				models.Error{
					Error: err.Error(),
				},
			)
			return
		}

		err = cache.Client.Set(ctx, fmt.Sprintf("scoreboard/check/%s", check), string(redisObject), config.Redis.FastRefresh).Err()
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				models.Error{
					Error: err.Error(),
				},
			)
			return
		}
	}

	ctx.JSON(http.StatusOK, checkScoreboard)
}

// CheckRound returns the scoreboard for a given check from a given round
//
// @Summary Get the scoreboard for a given check from a given round
// @Description Get the scoreboard for a given check from a given round
// @Tags scoreboard
// @Accept json
// @Produce json
// @Param check path string true "Check name"
// @Param round path string true "Round number"
// @Success 200 {object} structs.CheckScoreboard
// @Failure 500 {object} models.Error
// @Router /scoreboard/check/{check}/{round} [get]
func CheckRound(ctx *gin.Context) {
	check := ctx.Param("check")
	roundStr := ctx.Param("round")

	if config.Production {
		cachedData, err := cache.Client.Get(ctx, fmt.Sprintf("scoreboard/check/%s/%s", check, roundStr)).Result()
		if err != nil && err != redis.Nil {
			ctx.JSON(
				http.StatusInternalServerError,
				models.Error{
					Error: err.Error(),
				},
			)
			return
		}

		if err == nil {
			ctx.String(http.StatusOK, cachedData)
			return
		}
	}

	round, err := strconv.Atoi(roundStr)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	checkScoreboard, err := data.Scoreboard.CheckRound(check, round, 10)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	if config.Production {
		redisObject, err := json.Marshal(checkScoreboard)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				models.Error{
					Error: err.Error(),
				},
			)
			return
		}

		err = cache.Client.Set(ctx, fmt.Sprintf("scoreboard/check/%s/%s", check, roundStr), string(redisObject), config.Redis.FastRefresh).Err()
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				models.Error{
					Error: err.Error(),
				},
			)
			return
		}
	}

	ctx.JSON(http.StatusOK, checkScoreboard)
}
