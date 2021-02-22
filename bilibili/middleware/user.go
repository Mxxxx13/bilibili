package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err1 := c.Request.Cookie("username")
		if err1 == nil {
			c.Next()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "需要登录",
		})
		c.Abort()
	}
}
