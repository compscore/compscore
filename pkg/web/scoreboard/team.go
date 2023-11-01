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

// Scoreboard returns the scoreboard of a given team for the current round
//
// @Summary Get scoreboard of a given team for the current round
// @Description Get scoreboard of a given team for the current round
// @Tags scoreboard
// @Produce json
// @Param team path int true "Team number"
// @Success 200 {object} []structs.TeamScoreboard
// @Failure 500 {object} models.Error
// @Router /scoreboard/team/{team} [get]
func Team(ctx *gin.Context) {
	teamStr := ctx.Param("team")

	if config.Production {
		cachedData, err := cache.Client.Get(ctx, fmt.Sprintf("scoreboard/team/%s", teamStr)).Result()
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
			http.StatusBadRequest,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	teamScoreboard, err := data.Scoreboard.Team(team, 15)
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
		redisObject, err := json.Marshal(teamScoreboard)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				models.Error{
					Error: err.Error(),
				},
			)
			return
		}

		err = cache.Client.Set(ctx, fmt.Sprintf("scoreboard/team/%s", teamStr), string(redisObject), config.Redis.FastRefresh).Err()
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

	ctx.JSON(http.StatusOK, teamScoreboard)
}

// TeamRound returns the scoreboard of a given team for a given round
//
// @Summary Get scoreboard of a given team for a given round
// @Description Get scoreboard of a given team for a given round
// @Tags scoreboard
// @Produce json
// @Param team path int true "Team number"
// @Param round path int true "Round number"
// @Success 200 {object} []structs.TeamScoreboard
// @Failure 500 {object} models.Error
// @Router /scoreboard/team/{team}/{round} [get]
func TeamRound(ctx *gin.Context) {
	teamStr := ctx.Param("team")
	roundStr := ctx.Param("round")

	if config.Production {
		cachedData, err := cache.Client.Get(ctx, fmt.Sprintf("scoreboard/team/%s/%s", teamStr, roundStr)).Result()
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
			http.StatusBadRequest,
			models.Error{
				Error: err.Error(),
			},
		)
		return
	}

	round, err := strconv.Atoi(roundStr)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			models.Error{
				Error: err.Error(),
			},
		)

		return
	}

	teamScoreboard, err := data.Scoreboard.TeamRound(team, round, 15)
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
		redisObject, err := json.Marshal(teamScoreboard)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				models.Error{
					Error: err.Error(),
				},
			)
			return
		}

		err = cache.Client.Set(ctx, fmt.Sprintf("scoreboard/team/%s/%s", teamStr, roundStr), string(redisObject), config.Redis.FastRefresh).Err()
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

	ctx.JSON(http.StatusOK, teamScoreboard)
}
