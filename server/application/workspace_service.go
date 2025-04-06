package application

import (
	"server/domain"
	"server/infrastructure"
)

type WorkspaceService interface {
	ListWorkspaces() ([]domain.Workspace, error)
	GetWorkspace(id string) (*domain.WorkspaceWithChannels, error)
}

type workspaceServiceImpl struct {}

func NewWorkspaceService() WorkspaceService {
	return &workspaceServiceImpl{}
}

func (s *workspaceServiceImpl) ListWorkspaces() ([]domain.Workspace, error) {
	raw, err := infrastructure.FindAll()
	if err != nil {
		return nil, err
	}

	var result []domain.Workspace
	for _, w := range raw {
		result = append(result, domain.Workspace{
			ID:        w.ID,
			Name:      w.Name,
			Theme:     w.Theme,
			OwnerID:   w.OwnerID,
			CreatedAt: w.CreatedAt,
			UpdatedAt: w.UpdatedAt,
		})
	}
	return result, nil
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