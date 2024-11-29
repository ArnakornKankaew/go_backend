package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET METHOD",
		})
	})

	r.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST METHOD",
		})
	})

	r.PUT("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT METHOD",
		})
	})

	r.DELETE("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE METHOD",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
