package utils

import (
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// sinh token
func GetJwtToken(secretKey string, iat, seconds, userId int64, role int) (string, error) {
	claims := make(jwt.MapClaims)

	claims["exp"] = iat + seconds // thời gian hết hạn
	claims["iat"] = iat           //bắt đầu
	claims["userId"] = userId
	claims["role"] = role

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}

func GenerateResetToke() string {
	rand.Seed(time.Now().UnixNano())

	tokenLength := 20

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	token := make([]byte, tokenLength)

	for i := range token {
		token[i] = charset[rand.Intn(len(charset))]
	}
	return string(token)
}
