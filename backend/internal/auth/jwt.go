package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(
	userID int,
	email string,
	secret string,
) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString([]byte(secret))
}


func ValidateToken(
	tokenString string,
	secret string,
) (*jwt.Token, error) {

	return jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)
}