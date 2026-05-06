package router

import (
	"shrine-portal/backend/internal/handler"

	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo) {
	healthHandler := handler.NewHealthHandler()

	e.GET("/health", healthHandler.Show)
}
