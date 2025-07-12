package helpers

import (
	"fmt"
	"log"
	"time"

	"github.com/faisallbhr/gin-boilerplate/config"
	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret = []byte(config.GetEnv("JWT_SECRET", "secret"))

type CustomClaims struct {
	UserId uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userId uint) string {
	token, err := generateJwt(userId, time.Minute*15)
	if err != nil {
		log.Printf("Error generating access token: %v", err)
		return ""
	}
	return token
}

func GenerateRefreshToken(userId uint) string {
	token, err := generateJwt(userId, time.Hour*24*7)
	if err != nil {
		log.Printf("Error generating refresh token: %v", err)
		return ""
	}
	return token
}

func generateJwt(userId uint, duration time.Duration) (string, error) {
	expTime := time.Now().Add(duration)

	claims := CustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   fmt.Sprintf("%d", userId),
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}

func VerifyToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Invalid or expired token")
	}

	return claims, nil
}
