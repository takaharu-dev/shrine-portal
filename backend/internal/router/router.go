package router

import (
	"shrine-portal/backend/internal/handler"

	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo) {
	healthHandler := handler.NewHealthHandler()
	shrineHandler := handler.NewShrineHandler()

	e.GET("/health", healthHandler.Show)

	e.GET("/shrines", shrineHandler.Index)
	e.GET("/shrines/:id", shrineHandler.Show)
}
