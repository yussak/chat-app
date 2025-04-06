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