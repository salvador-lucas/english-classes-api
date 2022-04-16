package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/salvador-lucas/english-classes-api/models"
	"github.com/salvador-lucas/english-classes-api/routers"
	"github.com/salvador-lucas/english-classes-api/utils"
)

//TODO: obtain logger

func main() {
	// getting env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Panicln(err)
	}

	engine := gin.Default()

	//connect to database
	db, err := models.ObtainDbConnection()
	if err != nil {
		// panic(err)
		log.Panicln(err)
	}

	deps := utils.Dependencies{
		Db:     db,
		Logger: logger,
	}

	routers.InitializeRouters(engine, deps)

	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}
}
