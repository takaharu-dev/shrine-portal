package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ShrineHandler struct{}

func NewShrineHandler() *ShrineHandler {
	return &ShrineHandler{}
}

func (h *ShrineHandler) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, []string{
		"湯島天満宮",
		"明治神宮",
	})
}

func (h *ShrineHandler) Show(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, map[string]string{
		"id":   id,
		"name": "湯島天満宮",
	})
}
