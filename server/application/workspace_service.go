package application

import (
	"server/domain"
)

type WorkspaceService interface {
	ListWorkspaces() ([]domain.Workspace, error)
	GetWorkspace(id string) (*domain.WorkspaceWithChannels, error)
	ListSidebarProps() ([]domain.WorkspaceSidebarProps, error)
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

func (s *workspaceServiceImpl) ListSidebarProps() ([]domain.WorkspaceSidebarProps, error) {
	props, err := s.repo.GetSidebarProps()
	if err != nil {
		return nil, err
	}
	return props, nil
}