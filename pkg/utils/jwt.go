package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/yu1er/gin-blog/config"
)

var JWTSecret = []byte(config.JwtSecret)

type Claims struct {
	username string
	jwt.StandardClaims
}

// 生成jwt token, claims中附带username
func GenerateToken(username string) (string, error) {
	ddl := time.Now().Add(3 * time.Hour)

	claims := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: ddl.Unix(),
			Issuer:    "gin-blog",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString(JWTSecret)
	return token, err
}

// 解析jwt token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, err
		}
	}
	return nil, err
}
