package controllers

import (
	"net/http"
	"server/db"
	"server/models"

	"github.com/labstack/echo/v4"
)

func ListWorkspaces(c echo.Context) error {
	query := `
		SELECT
			w.id,
			w.name,
			w.owner_id,
			w.created_at,
			w.updated_at
		FROM workspaces w
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to find workspaces"})
	}
	defer rows.Close()
	
	var workspaces []models.Workspace
	for rows.Next() {
		var workspace models.Workspace
		if err := rows.Scan(&workspace.ID, &workspace.Name, &workspace.OwnerID, &workspace.CreatedAt, &workspace.UpdatedAt); err != nil {
			// if err := rows.Scan(&workspace.ID, &workspace.Name, &workspace.OwnerID, &workspace.CreatedAt, &workspace.UpdatedAt); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to scan workspace"})
		}
		workspaces = append(workspaces, workspace)
	}
	return c.JSON(http.StatusOK, workspaces)
}

func CreateWorkspace(c echo.Context) error {
	var req struct {
		Email       string `json:"email"`
		Name        string `json:"name"`
		DisplayName string `json:"displayName"`
	}

	// user existsは/users/existsでチェックしてるのでこっちでは不要そう？いや必要そう
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if req.Email == "" || req.Name == "" || req.DisplayName == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "メールアドレス、ワークスペース名、表示名は必須です",
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

	// ユーザーを検索
	user, err := models.FindUserByEmail(db.DB, req.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to find user"})
	}

	if user == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}

	ownerID := user.ID
	if ownerID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}

	workspace := models.Workspace{
		OwnerID: ownerID,
		Name:    req.Name,
	}

	query := `INSERT INTO workspaces (owner_id, name) VALUES ($1, $2) RETURNING id, created_at, updated_at`
	err = tx.QueryRow(query, workspace.OwnerID, req.Name).Scan(&workspace.ID, &workspace.CreatedAt, &workspace.UpdatedAt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "ワークスペース作成エラー",
		})
	}

	// ワークスペースメンバーを作成
	workspaceMember := models.WorkspaceMember{
		WorkspaceID: workspace.ID,
		UserID:      user.ID,
		DisplayName: req.DisplayName,
		ImageURL:    user.Image,
	}

	query = `INSERT INTO workspace_members (workspace_id, user_id, display_name, image_url) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`
	err = tx.QueryRow(query, workspaceMember.WorkspaceID, workspaceMember.UserID, workspaceMember.DisplayName, workspaceMember.ImageURL).Scan(&workspaceMember.ID, &workspaceMember.CreatedAt, &workspaceMember.UpdatedAt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "ワークスペースメンバー作成エラー",
		})
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "トランザクションコミットエラー",
		})
	}

	return c.JSON(http.StatusOK, workspace)
}

func GetWorkspace(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID is required"})
	}

	query := `
		SELECT 
			id,
			name,
			owner_id,
			created_at,
			updated_at
		FROM workspaces
		WHERE id = $1
	`
	var workspace models.Workspace
	err := db.DB.QueryRow(query, id).Scan(
		&workspace.ID,
		&workspace.Name,
		&workspace.OwnerID,
		&workspace.CreatedAt,
		&workspace.UpdatedAt,
	)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "workspace not found"})
	}

	return c.JSON(http.StatusOK, workspace)
}