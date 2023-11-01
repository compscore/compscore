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

// Status returns the status scoreboard for a given team and check
//
// @Summary Get status scoreboard for a given team and check
// @Description Get status scoreboard for a given team and check
// @Tags scoreboard
// @Produce json
// @Param check path string true "Check name"
// @Param team path int true "Team number"
// @Success 200 {object} []structs.Status
// @Failure 500 {object} models.Error
// @Router /scoreboard/status/{team}/{check} [get]
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
			ctx.String(http.StatusOK, cachedData)
			return
		}
	}

	team, err := strconv.Atoi(teamStr)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			models.Error{
				Error: "invalid team",
			},
		)
		return
	}

	statusHistory, err := data.Scoreboard.History(check, team, 10)
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
		redisObject, err := json.Marshal(statusHistory)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				models.Error{
					Error: err.Error(),
				},
			)
			return
		}

		err = cache.Client.Set(ctx, fmt.Sprintf("scoreboard/status/%s/%s", check, teamStr), string(redisObject), config.Redis.FastRefresh).Err()
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

	ctx.JSON(http.StatusOK, statusHistory)
}

// StatusRound returns the status scoreboard for a given team and check from a given round
//
// @Summary Get status scoreboard for a given team and check from a given round
// @Description Get status scoreboard for a given team and check from a given round
// @Tags scoreboard
// @Produce json
// @Param check path string true "Check name"
// @Param team path int true "Team number"
// @Param round path int true "Round number"
// @Success 200 {object} []structs.Status
// @Failure 500 {object} models.Error
// @Router /scoreboard/status/{team}/{check}/{round} [get]
func StatusRound(ctx *gin.Context) {
	check := ctx.Param("check")
	teamStr := ctx.Param("team")
	roundStr := ctx.Param("round")

	if config.Production {
		cachedData, err := cache.Client.Get(ctx, fmt.Sprintf("scoreboard/status/%s/%s/%s", check, teamStr, roundStr)).Result()
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

	team, err := strconv.Atoi(teamStr)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	round, err := strconv.Atoi(roundStr)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			models.Error{
				Error: "invalid round",
			},
		)
		return
	}

	statusHistory, err := data.Scoreboard.HistoryRound(check, team, round, 10)
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
		redisObject, err := json.Marshal(statusHistory)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				models.Error{
					Error: err.Error(),
				},
			)
			return
		}

		err = cache.Client.Set(ctx, fmt.Sprintf("scoreboard/status/%s/%s/%s", check, teamStr, roundStr), string(redisObject), config.Redis.FastRefresh).Err()
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

	ctx.JSON(http.StatusOK, statusHistory)
}
