package ui

import (
	"net/http"
	"server/application"

	"github.com/labstack/echo/v4"
)

func ListWorkspaces(c echo.Context) error {
	workspaces, err := application.ListWorkspaces()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to find workspaces"})
	}

	return c.JSON(http.StatusOK, workspaces)
}