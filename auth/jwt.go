package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func SignJWTClaims(claims jwt.Claims) (string, error) {
	mySigningKey := []byte("segredo-jwt")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(mySigningKey)
}

func ValidateJWT(jwtString string) (string, error) {
	mySigningKey := []byte("segredo-jwt")

	var userClaims jwt.RegisteredClaims
	token, err := jwt.ParseWithClaims(jwtString, &userClaims, func(t *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("Token invalido")
	}

	if userClaims.ExpiresAt.Before(time.Now()) {
		return "", fmt.Errorf("Token expirado")
	}

	// Subject == username
	return userClaims.Subject, nil
}
