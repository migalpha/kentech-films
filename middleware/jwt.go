package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	film "github.com/migalpha/kentech-films"
	"github.com/migalpha/kentech-films/jwt"
)

// Check if a request has token or permission needed
func CheckJWT(repo film.TokenProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		var errorJWT error
		var token string

		rToken := c.Request.Header["Authorization"]

		if len(rToken) < 1 {
			errorJWT = film.ErrMissingToken
		} else {
			token = rToken[0]
			splitToken := strings.Split(token, "Bearer")
			token = strings.TrimSpace(splitToken[1])

			claims, err := jwt.ParseToken(token)
			if err != nil {
				errorJWT = film.ErrAuthCheckTokenFail
			} else {
				if time.Now().Unix() > claims.ExpiresAt {
					errorJWT = film.ErrAuthCheckTokenTimeout
				} else {
					c.Set("user_id", claims.ID)
				}
			}

			signature := jwt.GetTokenSignature(rToken)
			isBlacklisted, err := repo.IsTokenBlacklisted(context.Background(), signature)
			if err != nil {
				errorJWT = film.ErrBlacklistCheckToken
			}
			if isBlacklisted {
				errorJWT = film.ErrBlacklistToken
			}
		}

		if errorJWT != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": errorJWT.Error(),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
