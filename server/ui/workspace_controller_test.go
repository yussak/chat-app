package ui

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"server/domain"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// モックサービス（WorkspaceServiceインターフェース実装）
type mockWorkspaceService struct {
	workspaces []domain.Workspace
	err        error
}

func (m *mockWorkspaceService) ListWorkspaces() ([]domain.Workspace, error) {
	return m.workspaces, m.err
}

func TestListWorkspaces_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/workspaces", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// ダミーデータ
	now := time.Now()
	dummyWorkspaces := []domain.Workspace{
		{
			ID:        1,
			Name:      "Test Workspace",
			Theme:     "キャンペーン準備",
			OwnerID:   100,
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	// モックサービスを生成し、テスト用のダミーデータを返すように設定
	mockService := &mockWorkspaceService{
		workspaces: dummyWorkspaces,
		err:        nil,
	}

	// コントローラにモックサービスを注入
	handler := NewWorkspaceController(mockService)

	// 実行
	err := handler.ListWorkspaces(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// レスポンスのJSONを検証
	var response []domain.Workspace
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, len(dummyWorkspaces), len(response))

	// 各フィールドを個別に比較（time.Timeは time.Equal を使う）
	for i, expected := range dummyWorkspaces {
		actual := response[i]
		assert.Equal(t, expected.ID, actual.ID)
		assert.Equal(t, expected.Name, actual.Name)
		assert.Equal(t, expected.Theme, actual.Theme)
		assert.Equal(t, expected.OwnerID, actual.OwnerID)
		assert.True(t, expected.CreatedAt.Equal(actual.CreatedAt), "CreatedAt mismatch")
		assert.True(t, expected.UpdatedAt.Equal(actual.UpdatedAt), "UpdatedAt mismatch")
	}
}
