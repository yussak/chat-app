package infrastructure

import (
	"server/db"
	"server/domain"
)

type WorkspaceRepository struct{}

func NewWorkspaceRepository() *WorkspaceRepository {
	return &WorkspaceRepository{}
}

func (r *WorkspaceRepository) FindAll() ([]domain.Workspace, error) {
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

func (r *WorkspaceRepository) FindById(id string) (*domain.WorkspaceWithChannels, error) {
	// todo: sqlはまとめて実行
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
