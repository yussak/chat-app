package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"server/db"
	"server/models"
	"server/test"
	"server/ui"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// メモ go test ./... -v
// todo: レイヤードに変えたのでテストも書き換え(不要かもしれない)
func TestListWorkspaces(t *testing.T) {
	// テストデータベースの接続
	testDB := test.GetTestDB(t)
	defer testDB.Close()

	// グローバルDBをテストDBに置き換え
	db.DB = testDB

	// テストデータの準備
	// まずユーザーを作成
	var userID int
	err := testDB.QueryRow(`
		INSERT INTO users (name, email, image)
		VALUES ('Test User', 'test@example.com', 'test.jpg')
		RETURNING id
	`).Scan(&userID)
	assert.NoError(t, err)

	// ワークスペースを作成
	_, err = testDB.Exec(`
		INSERT INTO workspaces (name, owner_id, theme, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
	`, "Test Workspace 1", userID, "秋のキャンペーン")
	assert.NoError(t, err)

	// テスト後にデータを消す
	defer func() {
		testDB.Exec("TRUNCATE TABLE workspaces CASCADE")
		testDB.Exec("TRUNCATE TABLE users CASCADE")
	}()

	// Echoの設定
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/workspaces", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// テストの実行
	err = ui.ListWorkspaces(c)

	// アサーション
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// レスポンスの検証
	var workspaces []models.Workspace
	err = json.Unmarshal(rec.Body.Bytes(), &workspaces)
	assert.NoError(t, err)
	assert.NotNil(t, workspaces)
	assert.Len(t, workspaces, 1)

	// レスポンスの内容を検証
	workspace := workspaces[0]
	assert.Equal(t, "Test Workspace 1", workspace.Name)
	assert.Equal(t, userID, workspace.OwnerID)
	assert.Equal(t, "秋のキャンペーン", workspace.Theme)
} 