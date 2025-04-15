package application

import (
	"server/domain"
	"server/infrastructure"
)

type WorkspaceService interface {
	ListWorkspaces() ([]domain.Workspace, error)
	GetWorkspace(id string) (*domain.WorkspaceWithChannels, error)
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

func ListSidebarProps() ([]infrastructure.WorkspaceSidebarProps, error) {
	raw, err := infrastructure.GetSidebarProps()
	if err != nil {
		return nil, err
	}
	return raw, nil
}

func (s *workspaceServiceImpl) GetWorkspace(id string) (*domain.WorkspaceWithChannels, error) {
	workspace, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	return workspace, nil
}