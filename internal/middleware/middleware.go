package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Mariano-SI/twitter-api/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		token = strings.TrimSpace(token)

		if token == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		userId, userName, err := jwt.ValidateToken(token, secretKey, true)

		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		c.Set("userId", userId)
		c.Set("userName", userName)
		c.Next()
	}
}

func AuthRefreshTokenMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		token = strings.TrimSpace(token)

		if token == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		userId, userName, err := jwt.ValidateToken(token, secretKey, false)

		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		c.Set("userId", userId)
		c.Set("userName", userName)
		c.Next()
	}
}
