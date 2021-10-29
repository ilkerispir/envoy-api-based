package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	var uuid = uuid.New().String()
	
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"uuid": uuid,
		})
	})

	r.Run(":8080")
}