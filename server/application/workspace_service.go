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
	raw, err := infrastructure.FindById(id)
	if err != nil {
		return nil, err
	}

	// todo:dto使うべきなら対応
	return &domain.WorkspaceWithChannels{
		Workspace: domain.Workspace{
			ID:        raw.Workspace.ID,
			Name:      raw.Workspace.Name,
			Theme:     raw.Workspace.Theme,
			OwnerID:   raw.Workspace.OwnerID,
			CreatedAt: raw.Workspace.CreatedAt,
			UpdatedAt: raw.Workspace.UpdatedAt,
		},
		Channels: func() []domain.Channel {
			var channels []domain.Channel
			for _, c := range raw.Channels {
				channels = append(channels, domain.Channel{
					ID:          c.ID,
					WorkspaceID: c.WorkspaceID,
					Name:        c.Name,
					IsPublic:    c.IsPublic,
					CreatedAt:   c.CreatedAt,
					UpdatedAt:   c.UpdatedAt,
				})
			}
			return channels
		}(),
	}, nil
}