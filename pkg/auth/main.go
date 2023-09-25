package auth

import (
	"time"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/structs"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(team string) (string, int, error) {
	expiration := time.Now().Add(time.Duration(config.Web.Timeout) * time.Hour)

	claims := &structs.Claims{
		Team: team,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(config.Web.JWTKey))

	return tokenStr, int(expiration.Unix()), err
}

func ParseJWT(tokenString string) (*jwt.Token, *structs.Claims, error) {
	claims := &structs.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Web.JWTKey), nil
	})
	return token, claims, err
}
