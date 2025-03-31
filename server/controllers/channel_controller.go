package controllers

import (
	"net/http"
	"server/db"
	"server/models"

	"github.com/labstack/echo/v4"
)

func GetChannel(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID is required"})
	}

	query := `
		SELECT 
			id,
			name,
			created_at,
			updated_at
		FROM channels
		WHERE id = $1
	`
	var channel models.Channel
	err := db.DB.QueryRow(query, id).Scan(
		&channel.ID,
		&channel.Name,
		&channel.CreatedAt,
		&channel.UpdatedAt,
	)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "channel not found"})
	}

	return c.JSON(http.StatusOK, channel)
}
