package middleware

import "github.com/gin-gonic/gin"

func DatabaseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("test", "hello")
		c.Next()
	}
}
