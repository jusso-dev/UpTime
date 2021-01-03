package main

import (
	"github.com/gin-gonic/gin"
)

//URL struct for POST request /ping-endpoint
type URL struct {
	ENDPOINT string `json:"endpoint" binding:"required"`
}

func main() {
	router := gin.Default()

	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.POST("/ping-endpoint", func(c *gin.Context) {

		var url URL
		c.BindJSON(&url)

		pingRes := pingURL(url.ENDPOINT)

		if pingRes.SUCCESS {
			c.JSON(200, pingRes)
		} else {
			c.JSON(500, pingRes)
		}
	})

	router.Run(":8008")
}
