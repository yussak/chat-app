package infrastructure

import (
	"database/sql"
	"server/db"
	"server/domain"
)

type WorkspaceRepository struct{}

func NewWorkspaceRepository() *WorkspaceRepository {
	return &WorkspaceRepository{}
}

type WorkspaceSidebarProps struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	YoungestChannelID int64  `json:"youngestChannelId"`
}

type WorkspaceRepositoryImpl struct {}

func NewWorkspaceRepositoryImpl() domain.WorkspaceRepository {
	return &WorkspaceRepositoryImpl{}
}

func (r *WorkspaceRepositoryImpl) FindAll() ([]domain.Workspace, error) {
	query := `SELECT id, name, owner_id, theme, created_at, updated_at FROM workspaces`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workspaces []domain.Workspace
	for rows.Next() {
		var workspace domain.Workspace
		if err := rows.Scan(&workspace.ID, &workspace.Name, &workspace.OwnerID, &workspace.Theme, &workspace.CreatedAt, &workspace.UpdatedAt); err != nil {
			return nil, err
		}
		workspaces = append(workspaces, workspace)
	}

	return workspaces, nil
}

func (r *WorkspaceRepositoryImpl) FindById(id string) (*domain.WorkspaceWithChannels, error) {
	// ワークスペース情報を取得
	workspaceQuery := `SELECT id, name, owner_id, theme, created_at, updated_at FROM workspaces WHERE id = $1`
	var workspace domain.Workspace
	err := db.DB.QueryRow(workspaceQuery, id).Scan(&workspace.ID, &workspace.Name, &workspace.OwnerID, &workspace.Theme, &workspace.CreatedAt, &workspace.UpdatedAt)
	if err != nil {
		return nil, err
	}

	// チャンネル情報を取得
	channelsQuery := `SELECT id, workspace_id, name, is_public, created_at, updated_at FROM channels WHERE workspace_id = $1`
	rows, err := db.DB.Query(channelsQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var channels []domain.Channel
	for rows.Next() {
		var channel domain.Channel
		err := rows.Scan(&channel.ID, &channel.WorkspaceID, &channel.Name, &channel.IsPublic, &channel.CreatedAt, &channel.UpdatedAt)
		if err != nil {
			return nil, err
		}
		channels = append(channels, channel)
	}

	// ワークスペースとチャンネル情報を結合
	response := domain.WorkspaceWithChannels{
		Workspace: workspace,
		Channels:  channels,
	}

	return &response, nil
}

func GetSidebarProps() ([]WorkspaceSidebarProps, error) {
	// ワークスペースとその最小チャンネルIDを取得するクエリ
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

	var workspaces []WorkspaceSidebarProps
	for rows.Next() {
		var workspace WorkspaceSidebarProps
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