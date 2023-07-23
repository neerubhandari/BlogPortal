package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const SecretKey = "secret"

// GenerateJwt creates a new JWT with the given issuer.
// The token uses the HS256 signing method and expires in 24 hours.
// It's then signed with a predefined SecretKey and returned as a string.

func GenerateJwt(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	return claims.SignedString([]byte(SecretKey))
}

// ParseJwt decodes a JWT from a string, validates it with the predefined SecretKey,
// and returns the issuer if the token is valid. It returns an error if parsing or validation fails.

func ParseJwt(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}
	claims := token.Claims.(*jwt.StandardClaims)
	return claims.Issuer, nil
}
