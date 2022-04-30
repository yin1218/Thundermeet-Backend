package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

type authClaims struct {
	jwt.StandardClaims
	UserID string `json:"userId"`
}

// GenToken Create a new token
func GenToken(userId string) (string, error) {
	// load secret key
	jwtKeyString := os.Getenv("JWT_SECRET")
	jwtKey := []byte(jwtKeyString)

	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, authClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		UserID: userId,
	})
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func validateToken(tokenString string) (string, error) {
	jwtKeyString := os.Getenv("JWT_SECRET")
	jwtKey := []byte(jwtKeyString)
	var claims authClaims

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", errors.New("invalid token")
	}
	id := claims.UserID
	return id, nil
}
