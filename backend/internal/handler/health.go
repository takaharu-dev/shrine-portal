package handler

import (
	"net/http"

	"shrine-portal/backend/internal/response"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Show(c echo.Context) error {
	return c.JSON(http.StatusOK, response.HealthResponse{Status: "ok"})
}
