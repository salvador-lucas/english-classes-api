package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salvador-lucas/english-classes-api/services"
	"github.com/salvador-lucas/english-classes-api/utils"
)

type HealthController struct {
	ServiceFactory func() services.HealthService
}

// HealthCheck ...
// @Description Performs Health Check
// @Success 201 {object} views.BaseResponse
// @Failure 503 {object} views.BaseResponse
// @Router /health [get]
func (h *HealthController) HealthCheck(ctx *gin.Context) {
	result, errs := h.ServiceFactory().HealthCheck()
	if len(errs) != 0 {
		utils.SetResponse(ctx, http.StatusServiceUnavailable, result)
		return
	}
	utils.SetResponse(ctx, http.StatusOK, result)
}
