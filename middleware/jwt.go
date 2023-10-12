package middleware

import (
	"Waymon_api/pkg/e"
	"Waymon_api/pkg/waymon"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = e.Success
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.TokenError
		} else {
			claims, err := waymon.ParseToken(token)
			if err != nil {
				code = e.TokenError
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.TokenError
			}
		}
		if code != e.Success {
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
