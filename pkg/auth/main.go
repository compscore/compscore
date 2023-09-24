package auth

import (
	"fmt"
	"time"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/data"
	"github.com/compscore/compscore/pkg/ent"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Team string
	jwt.RegisteredClaims
}
type User struct {
	Username   string
	Expiration int
}

func GenerateJWT(team string) (string, int, error) {
	expiration := time.Now().Add(time.Duration(config.Web.Timeout) * time.Hour)

	claims := &Claims{
		Team: team,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(config.Web.JWTKey))

	return tokenStr, int(expiration.Unix()), err
}

func ParseJWT(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Web.JWTKey), nil
	})
	return token, claims, err
}

func Parse(ctx *gin.Context) (*ent.Team, error) {
	tokenString, err := ctx.Cookie("auth")
	if err != nil {
		return nil, fmt.Errorf("no auth cookie found")
	}

	_, claims, err := ParseJWT(tokenString)
	if err != nil {
		return nil, err
	}

	return data.Team.GetByName(claims.Team)
}
