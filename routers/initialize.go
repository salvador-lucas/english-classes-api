package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/salvador-lucas/english-classes-api/utils"
)

func InitializeRouters(engine *gin.Engine, deps utils.Dependencies) {

	//health
	HealthRoutes(engine, deps)
	//students
	StudentsRoutes(engine, deps)
	//students
	// PaymentsRoutes(engine, deps)
}
