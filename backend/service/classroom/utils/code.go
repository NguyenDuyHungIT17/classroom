package utils

import (
	"math/rand"
	"time"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZqwertyuiopasdfghjklzxcvbnm0123456789"

func GenerateClassCode() string {
	rand.Seed(time.Now().UnixNano())

	codeLength := 6
	result := make([]byte, codeLength)

	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}
