package handler

import (
	"net/http"

	"shrine-portal/backend/internal/service"

	"github.com/labstack/echo/v4"
)

type BenefitHandler struct {
	service *service.BenefitService
}

func NewBenefitHandler() *BenefitHandler {
	return &BenefitHandler{
		service: service.NewBenefitService(),
	}
}

func (h *BenefitHandler) Index(c echo.Context) error {
	benefits := h.service.FindAll()

	return c.JSON(http.StatusOK, benefits)
}
