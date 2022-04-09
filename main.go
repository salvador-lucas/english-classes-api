package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/salvador-lucas/english-classes-api/models"
)

func main() {
	r := gin.Default()

	// _, err := models.ObtainDbConnection()
	// if err != nil {
	// 	panic(err)
	// }

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.Run()
}
