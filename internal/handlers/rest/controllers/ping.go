package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPing(c *gin.Context) {
	test := c.MustGet("test")
	c.JSON(http.StatusOK, gin.H{
		"message": test,
	})
}
