package helpers

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateToken(id int, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expired in 24 hours
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
