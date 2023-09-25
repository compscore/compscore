package auth

import (
	"context"
	"fmt"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/data"
	"github.com/compscore/compscore/pkg/ent"
	"github.com/compscore/compscore/pkg/structs"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("auth")
	if err != nil {
		ctx.Next()
		return
	}

	claims := &structs.Claims{}
	jwtToken, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Web.JWTKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.Next()
			return
		}
		ctx.Next()
		return
	}
	if !jwtToken.Valid {
		ctx.Next()
		return
	}

	entTeam, err := data.Team.GetByName(claims.Team)
	if err != nil {
		ctx.Next()
		return
	}

	ctx.Request = ctx.Request.WithContext(
		context.WithValue(
			ctx.Request.Context(),
			structs.TEAM_CTX_KEY,
			entTeam,
		),
	)

	ctx.Next()
}

func Parse(ctx *gin.Context) (*ent.Team, error) {
	team, ok := ctx.Request.Context().Value(structs.TEAM_CTX_KEY).(*ent.Team)
	if !ok {
		return nil, fmt.Errorf("no team found in context")
	}
	return team, nil
}
