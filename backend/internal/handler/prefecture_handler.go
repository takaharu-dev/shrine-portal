package handler

import (
	"net/http"

	"shrine-portal/backend/internal/service"

	"github.com/labstack/echo/v4"
)

type PrefectureHandler struct {
	service *service.PrefectureService
}

func NewPrefectureHandler() *PrefectureHandler {
	return &PrefectureHandler{
		service: service.NewPrefectureService(),
	}
}

func (h *PrefectureHandler) Index(c echo.Context) error {
	prefectures := h.service.FindAll()

	return c.JSON(http.StatusOK, prefectures)
}
