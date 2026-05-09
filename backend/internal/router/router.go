package router

import (
	"shrine-portal/backend/internal/handler"

	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo) {
	healthHandler := handler.NewHealthHandler()
	shrineHandler := handler.NewShrineHandler()
	prefectureHandler := handler.NewPrefectureHandler()
	benefitHandler := handler.NewBenefitHandler()

	e.GET("/health", healthHandler.Show)

	e.GET("/shrines", shrineHandler.Index)
	e.GET("/shrines/:id", shrineHandler.Show)

	e.GET("/prefectures", prefectureHandler.Index)

	e.GET("/benefits", benefitHandler.Index)
}
