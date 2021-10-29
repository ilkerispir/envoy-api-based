package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"github.com/google/uuid"
)

func main() {
	var uuid = uuid.New().String()
	
	r := gin.Default()
	
	r.Use(favicon.New("./favicon.ico"))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"uuid": uuid,
		})
	})

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"message": "ok",
		})
	})

	r.Run(":8081")
}