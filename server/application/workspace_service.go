package application

import (
	"server/domain"
	"server/infrastructure"
)

func ListWorkspaces() ([]domain.Workspace, error) {
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