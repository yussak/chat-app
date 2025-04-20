package ui

import (
	"net/http"
	"server/application"
	"server/db"

	"github.com/labstack/echo/v4"
)

type WorkspaceController struct {
	Service application.WorkspaceService
}

func NewWorkspaceController(s application.WorkspaceService) *WorkspaceController {
	return &WorkspaceController{Service: s}
}

func (h *WorkspaceController) ListWorkspacesHandler(c echo.Context) error {
	workspaces, err := h.Service.ListWorkspaces()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to find workspaces"})
	}

	return c.JSON(http.StatusOK, workspaces)
}

func (h *WorkspaceController) GetWorkspace(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID is required"})
	}

	workspace, err := h.Service.GetWorkspace(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get workspace"})
	}

	return c.JSON(http.StatusOK, workspace)
}

func (h *WorkspaceController) CreateWorkspaceHandler(c echo.Context) error {
	var req struct {
		Email       string `json:"email"`
		Name        string `json:"name"`
		DisplayName string `json:"displayName"`
		Theme       string `json:"theme"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if req.Email == "" || req.Name == "" || req.DisplayName == "" || req.Theme == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "メールアドレス、ワークスペース名、表示名、テーマは必須です",
		})
	}

	// トランザクション開始
	tx, err := db.DB.Begin()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "トランザクション開始エラー",
		})
	}
	defer tx.Rollback()

	err = h.Service.CreateWorkspace(tx, req.DisplayName, req.Name, req.Theme, req.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	if err := tx.Commit(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "トランザクションコミットエラー",
		})
	}

	return nil
}
