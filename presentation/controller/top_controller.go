package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Top() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	}
}