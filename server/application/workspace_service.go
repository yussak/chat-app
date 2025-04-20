package application

import (
	"database/sql"
	"server/db"
	"server/domain"
	"server/models"
)

type WorkspaceService interface {
	ListWorkspaces() ([]domain.Workspace, error)
	GetWorkspace(id string) (*domain.WorkspaceWithChannels, error)
	CreateWorkspace(tx *sql.Tx, displayName string, name string, theme string, email string) error
}

type workspaceServiceImpl struct {
	repo domain.WorkspaceRepository
}

func NewWorkspaceService(repo domain.WorkspaceRepository) WorkspaceService {
	return &workspaceServiceImpl{repo: repo}
}

func (s *workspaceServiceImpl) ListWorkspaces() ([]domain.Workspace, error) {
	workspaces, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return workspaces, nil
}

func (s *workspaceServiceImpl) GetWorkspace(id string) (*domain.WorkspaceWithChannels, error) {
	workspace, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	return workspace, nil
}

func (s *workspaceServiceImpl) CreateWorkspace(tx *sql.Tx, displayName string, name string, theme string, email string) error {
	// ユーザーを検索
	// todo:models依存を廃止
	user, err := models.FindUserByEmail(db.DB, email)
	if err != nil {
		return err
	}

	if user == nil {
		return err
	}

	return s.repo.CreateWorkspace(tx, displayName, name, theme, user)
}
