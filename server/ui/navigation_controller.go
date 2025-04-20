package ui

import (
	"net/http"
	"server/application"

	"github.com/labstack/echo/v4"
)

type NavigationController struct {
	Service application.NavigationService
}

func NewNavigationController(s application.NavigationService) *NavigationController {
	return &NavigationController{Service: s}
}

func (h *NavigationController) GetSidebarProps(c echo.Context) error {
	workspaceNavItems, err := h.Service.ListSidebarProps()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to find sidebar workspaceNavItems"})
	}

	return c.JSON(http.StatusOK, workspaceNavItems)
}
