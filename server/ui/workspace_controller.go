package ui

import (
	"net/http"
	"server/application"

	"github.com/labstack/echo/v4"
)

type WorkspaceController struct {
	Service application.WorkspaceService
}

func NewWorkspaceController(s application.WorkspaceService) *WorkspaceController {
	return &WorkspaceController{Service: s}
}

func (h *WorkspaceController) ListWorkspaces(c echo.Context) error {
	workspaces, err := h.Service.ListWorkspaces()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to find workspaces"})
	}

	return c.JSON(http.StatusOK, workspaces)
}

func GetWorkspace(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID is required"})
	}

	workspace, err := application.GetWorkspace(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get workspace"})
	}

	return c.JSON(http.StatusOK, workspace)
}

func GetSidebarProps(c echo.Context) error {
	workspaces, err := application.ListSidebarProps()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to find workspaces"})
	}

	return c.JSON(http.StatusOK, workspaces)
}