package controllers

import (
	"net/http"
	"server/models"

	"github.com/labstack/echo/v4"
)

func GetChannel(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID is required"})
	}

	channel, err := models.GetChannel(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "channel not found"})
	}

	return c.JSON(http.StatusOK, channel)
}
