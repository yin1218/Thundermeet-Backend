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
	UserID int64 `json:"userId"`
}

// GenToken Create a new token
func GenToken(userId int64) (string, error) {
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

func validateToken(tokenString string) (int64, error) {
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
		return 0, err
	}
	if !token.Valid {
		return 0, errors.New("invalid token")
	}
	id := claims.UserID
	return id, nil
}
