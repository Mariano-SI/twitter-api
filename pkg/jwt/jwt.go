package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userId, username, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       userId,
		"username": username,
		"exp":      time.Now().Add(60 * time.Minute).Unix(),
	})

	key := []byte(secretKey)

	tokenStr, err := token.SignedString(key)

	return tokenStr, err
}

func ValidateToken(tokenStr, secretKey string, withClaimValidation bool) (string, string, error) {
	var (
		key    = []byte(secretKey)
		claims = jwt.MapClaims{}
		token  *jwt.Token
		err    error
	)

	if withClaimValidation {
		token, err = jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return key, nil
		})
	} else {
		token, err = jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return key, nil
		}, jwt.WithoutClaimsValidation())
	}

	if err != nil {
		return "", "", err
	}

	if !token.Valid {
		return "", "", errors.New("invalid token")
	}

	userId, userName := claims["id"].(string), claims["username"].(string)

	return userId, userName, nil
}
