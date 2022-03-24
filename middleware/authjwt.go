package middleware

import (
	"ProjectSavePassword/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func MiddlewareJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		bearerToken := tokenString[7:]
		token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("projectinidilindungi"), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, helper.ResponseAPI("Error when parsing token", http.StatusNotAcceptable, err.Error()))
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId := uint(claims["id"].(float64))
			c.Set("UserLoggedIn", userId)
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.ResponseAPI("Error when validate token", http.StatusUnauthorized, err.Error()))
		}

	}
}
