package main

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/salvador-lucas/english-classes-api/models"
	"github.com/salvador-lucas/english-classes-api/routers"
	"github.com/salvador-lucas/english-classes-api/utils"
)

func obtainLogger() (*log.Logger, error) {
	logger := log.New()
	logger.SetOutput(os.Stdout)

	logLevel := os.Getenv("LOG_LEVEL")

	switch logLevel {
	case "INFO":
		logger.SetLevel(log.InfoLevel)
	case "ERROR":
		logger.SetLevel(log.InfoLevel)
	case "DEBUG":
		logger.SetLevel(log.InfoLevel)
	default:
		return nil, errors.New("Invalid log level provided")
	}
	return logger, nil
}

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
	logger, err := obtainLogger()
	if err != nil {
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
