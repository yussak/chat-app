package controllers

import (
	"fmt"
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
		Email string `json:"email"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if req.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "email is required"})
	}

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
		// Email:   req.Email,
		OwnerID: ownerID,
	}

	query := `INSERT INTO workspaces (owner_id, name) VALUES ($1, $2) RETURNING id, created_at, updated_at`
	err = db.DB.QueryRow(query, workspace.OwnerID, req.Email).Scan(&workspace.ID, &workspace.CreatedAt, &workspace.UpdatedAt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create workspace"})
	}

	return c.JSON(http.StatusOK, workspace)
}