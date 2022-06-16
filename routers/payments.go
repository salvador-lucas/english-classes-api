package routers

// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/salvador-lucas/english-classes-api/controllers"
// 	"github.com/salvador-lucas/english-classes-api/services"
// 	"github.com/salvador-lucas/english-classes-api/utils"
// )

// func PaymentsRoutes(engine *gin.Engine, deps utils.Dependencies) {
// 	paymentsController := controllers.PaymentsController{
// 		ServiceFactory: func() services.PaymentsService {
// 			return services.NewPaymentsService(deps)
// 		},
// 	}

// 	paymentsGroup := engine.Group("payments")
// 	// studenGroup.GET("/", paymentsController.GetAllPayments)
// 	paymentsGroup.POST("/:studentId", paymentsController.AddPayment)
// }
