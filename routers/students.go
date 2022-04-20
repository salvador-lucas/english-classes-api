package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/salvador-lucas/english-classes-api/controllers"
	"github.com/salvador-lucas/english-classes-api/services"
	"github.com/salvador-lucas/english-classes-api/utils"
)

func StudentsRoutes(engine *gin.Engine, deps utils.Dependencies) {
	studentController := controllers.StudentController{
		ServiceFactory: func() services.StudentsService {
			return services.NewStudentService(deps)
		},
	}

	studenGroup := engine.Group("students")
	studenGroup.GET("/", studentController.GetAllStudents)
	studenGroup.POST("/", studentController.AddStudent)
}
