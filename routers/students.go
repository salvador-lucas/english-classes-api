package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/salvador-lucas/english-classes-api/controllers"
	"github.com/salvador-lucas/english-classes-api/services"
	"github.com/salvador-lucas/english-classes-api/utils"
)

func StudentsRoutes(engine *gin.Engine, deps utils.Dependencies) {
	studentController := controllers.StudentsController{
		ServiceFactory: func() services.StudentsService {
			return services.NewStudentService(deps)
		},
	}

	studentsGroup := engine.Group("students")
	studentsGroup.GET("/", studentController.GetAllStudents)
	studentsGroup.POST("/", studentController.AddStudent)
	studentsGroup.POST("/:id/payment", studentController.AddPaymentStudent)
	studentsGroup.GET("/:id/payment", studentController.GetStudentPayments)
}
