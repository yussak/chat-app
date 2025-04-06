package controllers

import (
	"net/http"
	"server/db"
	"server/models"

	"github.com/labstack/echo/v4"
)


func CreateWorkspace(c echo.Context) error {
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

	// ユーザーを検索
	user, err := models.FindUserByEmail(db.DB, req.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to find user"})
	}

	if user == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}

	workspace := models.Workspace{
		OwnerID: user.ID,
		Name:    req.Name,
		Theme:   req.Theme,
	}

	if err := models.CreateWorkspace(tx, &workspace, req.DisplayName, user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
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