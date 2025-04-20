package ui

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"server/application"
	"server/domain"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// モックリポジトリ（WorkspaceRepositoryインターフェース実装）
type mockWorkspaceRepository struct {
	workspaces []domain.Workspace
	workspace  *domain.WorkspaceWithChannels
	err        error
}

func (m *mockWorkspaceRepository) FindAll() ([]domain.Workspace, error) {
	return m.workspaces, m.err
}

func (m *mockWorkspaceRepository) FindById(id string) (*domain.WorkspaceWithChannels, error) {
	return m.workspace, m.err
}

// モックサービス（WorkspaceServiceインターフェース実装）
type mockWorkspaceService struct {
	workspaces []domain.Workspace
	workspace  *domain.WorkspaceWithChannels
	err        error
}

func (m *mockWorkspaceService) ListWorkspaces() ([]domain.Workspace, error) {
	return m.workspaces, m.err
}

func (m *mockWorkspaceService) GetWorkspace(id string) (*domain.WorkspaceWithChannels, error) {
	return m.workspace, m.err
}

// todo:他の文も書く
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

	// モックリポジトリを生成し、テスト用のダミーデータを返すように設定
	mockRepo := &mockWorkspaceRepository{
		workspaces: dummyWorkspaces,
		err:        nil,
	}

	// ワークスペースサービスにモックリポジトリを注入
	service := application.NewWorkspaceService(mockRepo)

	// コントローラにサービスを注入
	handler := NewWorkspaceController(service)

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

func TestGetWorkspace_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/workspaces/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/workspaces/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// ダミーデータ
	now := time.Now()
	dummyWorkspace := &domain.WorkspaceWithChannels{
		Workspace: domain.Workspace{
			ID:        1,
			Name:      "Test Workspace",
			Theme:     "キャンペーン準備",
			OwnerID:   100,
			CreatedAt: now,
			UpdatedAt: now,
		},
		Channels: []domain.Channel{
			{
				ID:          1,
				WorkspaceID: 1,
				Name:        "general",
				IsPublic:    true,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
		},
	}

	// モックサービスを生成し、テスト用のダミーデータを返すように設定
	mockService := &mockWorkspaceService{
		workspace: dummyWorkspace,
		err:       nil,
	}

	// コントローラにモックサービスを注入
	handler := NewWorkspaceController(mockService)

	// 実行
	err := handler.GetWorkspace(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// レスポンスのJSONを検証
	var response domain.WorkspaceWithChannels
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, dummyWorkspace.ID, response.ID)
	assert.Equal(t, dummyWorkspace.Name, response.Name)
	assert.Equal(t, dummyWorkspace.Theme, response.Theme)
	assert.Equal(t, dummyWorkspace.OwnerID, response.OwnerID)
	assert.True(t, dummyWorkspace.CreatedAt.Equal(response.CreatedAt), "CreatedAt mismatch")
	assert.True(t, dummyWorkspace.UpdatedAt.Equal(response.UpdatedAt), "UpdatedAt mismatch")
	assert.Equal(t, len(dummyWorkspace.Channels), len(response.Channels))
}
