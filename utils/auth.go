package utils

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID string, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("your_secret_key"))
}

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("your_secret_key"), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, http.ErrNoCookie
} 