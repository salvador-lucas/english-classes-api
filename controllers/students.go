package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salvador-lucas/english-classes-api/services"
	"github.com/salvador-lucas/english-classes-api/utils"
)

type StudentController struct {
	ServiceFactory func() services.StudentsService
}

// GetAllStudents ...
// @Description Get all students from database
// @Param message body views.AddSolarSystemRequest false "Add Solar System Request"
// @Success 200 {object} views.BaseResponse
// @Failure 400 {object} views.BaseResponse
// @Failure 500 {object} views.BaseResponse
// @Router /students [get]
func (s *StudentController) GetAllStudents(ctx *gin.Context) {
	result, err := s.ServiceFactory().GetAllStudents()
	if err != nil {
		utils.SetResponse(ctx, utils.GetStatusErrorCode(err), err)
		return
	}

	utils.SetResponse(ctx, http.StatusOK, result)
}
