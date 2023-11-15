package controller

import (
	"net/http"

	"enigmacamp.com/be-lms-university/model/dto"
	"enigmacamp.com/be-lms-university/usecase"
	"github.com/gin-gonic/gin"
)

type EnrollmentController struct {
	uc usecase.EnrollmentUseCase
	rg *gin.RouterGroup
}

func (e *EnrollmentController) createHandler(ctx *gin.Context) {
	var payload dto.EnrollmentRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	payloadResponse, err := e.uc.RegisterNewEnrollment(payload)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
	}

	dto.SendSingleResponse(ctx, http.StatusCreated, "Ok", payloadResponse)
}

func (e *EnrollmentController) Route() {
	e.rg.POST("/enrollments", e.createHandler)
}

func NewEnrollmentController(uc usecase.EnrollmentUseCase, rg *gin.RouterGroup) *EnrollmentController {
	return &EnrollmentController{uc: uc, rg: rg}
}
