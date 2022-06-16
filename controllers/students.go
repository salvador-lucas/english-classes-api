package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/salvador-lucas/english-classes-api/services"
	"github.com/salvador-lucas/english-classes-api/utils"
	"github.com/salvador-lucas/english-classes-api/views"
)

type StudentsController struct {
	ServiceFactory func() services.StudentsService
}

// GetAllStudents ...
// @Description Get all students from database
// @Param message body views.GetAllStudentsResponse false "Get all students"
// @Success 200 {object} views.BaseResponse
// @Failure 400 {object} views.BaseResponse
// @Failure 500 {object} views.BaseResponse
// @Router /students [get]
func (s *StudentsController) GetAllStudents(ctx *gin.Context) {
	result, err := s.ServiceFactory().GetAllStudents()
	if err != nil {
		utils.SetResponse(ctx, utils.GetStatusErrorCode(err), err)
		return
	}

	utils.SetResponse(ctx, http.StatusOK, result)
}

// AddStudent ...
// @Insert a new student in the database
// @Param message body views.AddStudentResponse false "Create new student"
// @Created 201 {object} views.BaseResponse
// @Failure 400 {object} views.BaseResponse
// @Failure 500 {object} views.BaseResponse
// @Router /students [post]
func (s *StudentsController) AddStudent(ctx *gin.Context) {
	var request views.AddStudentRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.SetResponse(ctx, http.StatusBadRequest, err)
		return
	}

	result, err := s.ServiceFactory().AddStudent(&request)
	if err != nil {
		utils.SetResponse(ctx, utils.GetStatusErrorCode(err), err)
		return
	}
	utils.SetResponse(ctx, http.StatusCreated, result)
}

// AddPaymentStudent ...
// @Insert a new payment for a specific student in the database
// @Param message body views.AddPaymentRequest false "Create new payment student"
// @Created 201 {object} views.BaseResponse
// @Failure 400 {object} views.BaseResponse
// @Failure 500 {object} views.BaseResponse
// @Router /students [post]
func (s *StudentsController) AddPaymentStudent(ctx *gin.Context) {
	studentID, errParam := strconv.Atoi(ctx.Param("id"))
	if errParam != nil {
		utils.SetResponse(ctx, http.StatusBadRequest, errParam)
		return
	}

	var request views.AddPaymentRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.SetResponse(ctx, http.StatusBadRequest, err)
		return
	}
	result, err := s.ServiceFactory().AddStudentPayment(studentID, &request)
	if err != nil {
		utils.SetResponse(ctx, utils.GetStatusErrorCode(err), err)
		return
	}
	utils.SetResponse(ctx, http.StatusCreated, result)
}
