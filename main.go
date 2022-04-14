package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/salvador-lucas/english-classes-api/models"
	// "github.com/salvador-lucas/english-classes-api/models"
)

func main() {
	// getting env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Panicln(err)
	}

	r := gin.Default()

	//connect to database
	_, err = models.ObtainDbConnection()
	if err != nil {
		// panic(err)
		log.Panicln(err)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.Run()
}
