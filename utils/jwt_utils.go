package utils

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(email string) (string, int64, error) {
	JWTSecretKey := os.Getenv("JWT_SECRET_KEY")
	JWTDUration := os.Getenv("JWT_DURATION")
	durationHours := 8
	if parsedDuration, err := strconv.Atoi(JWTDUration); err == nil && parsedDuration > 0 {
		durationHours = parsedDuration
	}

	expirationTime := time.Now().Add(time.Duration(durationHours) * time.Hour).Unix()

	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWTSecretKey))
	if err != nil {
		return "", 0, err
	}

	return tokenString, expirationTime, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	JWTSecretKey := os.Getenv("JWT_SECRET_KEY")

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
