package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Protect(signature []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		tokenz := strings.TrimPrefix(authorization, "Bearer ")
		
		_, err := jwt.Parse(tokenz, func(token *jwt.Token) (interface{}, error) {
			//เช็ค Method ว่าเป็น MethodHMAC ไหม
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				c.JSON(http.StatusConflict, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])) 
			}
			//เช็ค Method ผ่าน return signature ให้
			return []byte(signature), nil
		})
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
		}
		c.Next()
	}
}