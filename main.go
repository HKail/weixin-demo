package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "Hello world!"})
	})

	log.Fatalln(router.Run(":80"))
}
