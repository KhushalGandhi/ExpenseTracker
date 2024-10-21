package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

func GenerateJWT(userID uint) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // Token valid for 72 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	secretKey := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
