package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/salvador-lucas/english-classes-api/controllers"
	"github.com/salvador-lucas/english-classes-api/services"
	"github.com/salvador-lucas/english-classes-api/utils"
)

func HealthRoutes(engine *gin.Engine, deps utils.Dependencies) {
	healthController := controllers.HealthController{
		ServiceFactory: func() services.HealthService {
			return services.NewHealthService(deps)
		},
	}

	healthGroup := engine.Group("health")
	healthGroup.GET("/", healthController.HealthCheck)
}
