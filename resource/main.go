package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"github.com/google/uuid"
	"os"
)

func main() {
	var uuid = uuid.New().String()
	var hostname,_ = os.Hostname()
	
	r := gin.Default()
	
	r.Use(favicon.New("./favicon.ico"))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"uuid": uuid,
			"hostname": hostname,
		})
	})

	r.Run(":8080")
}