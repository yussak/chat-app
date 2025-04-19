package infrastructure

import (
	"database/sql"
	"server/db"
	"server/domain"
)

type NavigationRepository struct{}

func NewNavigationRepository() *NavigationRepository {
	return &NavigationRepository{}
}

func (r *NavigationRepository) GetSidebarProps() ([]domain.NavigationSidebarProps, error) {
	query := `
		SELECT w.id, w.name, MIN(c.id) as channel_id
		FROM workspaces w
		LEFT JOIN channels c ON w.id = c.workspace_id
		GROUP BY w.id, w.name
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workspaces []domain.NavigationSidebarProps
	for rows.Next() {
		var workspace domain.NavigationSidebarProps
		var channelID sql.NullInt64
		if err := rows.Scan(&workspace.ID, &workspace.Name, &channelID); err != nil {
			return nil, err
		}
		if channelID.Valid {
			workspace.YoungestChannelID = channelID.Int64
		}
		workspaces = append(workspaces, workspace)
	}

	return workspaces, nil
}
