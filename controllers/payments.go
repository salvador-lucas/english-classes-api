package controllers

// import (
// 	"github.com/salvador-lucas/english-classes-api/services"
// )

// type PaymentsController struct {
// 	ServiceFactory func() services.PaymentsService
// }

// GetAllPayments ...
// @Description Get all students from database
// @Param message body views.GetAllPaymentsResponse "Get all students"
// @Success 200 {object} views.BaseResponse
// @Failure 400 {object} views.BaseResponse
// @Failure 500 {object} views.BaseResponse
// @Router /students [get]
// func (p *PaymentsController) GetAllPayments(ctx *gin.Context) {
// 	result, err := p.ServiceFactory().GetAllPayments()
// 	if err != nil {
// 		utils.SetResponse(ctx, utils.GetStatusErrorCode(err), err)
// 		return
// 	}

// 	utils.SetResponse(ctx, http.StatusOK, result)
// }

// AddPayment ...
// @Insert a new student in the database
// @Param message body views.AddPaymentResponse false "Create new student"
// @Created 201 {object} views.BaseResponse
// @Failure 400 {object} views.BaseResponse
// @Failure 500 {object} views.BaseResponse
// @Router /students [post]
// func (p *PaymentsController) AddPayment(ctx *gin.Context) {
// 	var request views.AddPaymentRequest
// 	if err := ctx.ShouldBindJSON(&request); err != nil {
// 		utils.SetResponse(ctx, http.StatusBadRequest, err)
// 		return
// 	}

// 	result, err := p.ServiceFactory().AddPayment(&request)
// 	if err != nil {
// 		utils.SetResponse(ctx, utils.GetStatusErrorCode(err), err)
// 		return
// 	}
// 	utils.SetResponse(ctx, http.StatusCreated, result)
// }
