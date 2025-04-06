package infrastructure

import (
	"server/db"
	"time"
)

type Workspace struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	OwnerID   int       `json:"ownerId"`
	Theme     string    `json:"theme"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type WorkspaceWithChannels struct {
	Workspace
	Channels []Channel `json:"channels"`
}

type Channel struct {
	ID          int64     `json:"id"`
	WorkspaceID int64     `json:"workspace_id"`
	Name        string    `json:"name"`
	IsPublic    bool      `json:"is_public"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FindAll() ([]Workspace, error) {
	query := `SELECT id, name, owner_id, theme, created_at, updated_at FROM workspaces`
		rows, err := db.DB.Query(query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var workspaces []Workspace
		for rows.Next() {
			var workspace Workspace
			if err := rows.Scan(&workspace.ID, &workspace.Name, &workspace.OwnerID, &workspace.Theme, &workspace.CreatedAt, &workspace.UpdatedAt); err != nil {
				return nil, err
			}
			workspaces = append(workspaces, workspace)
		}

		return workspaces, nil
}

func FindById(id string) (*WorkspaceWithChannels, error) {
	// ワークスペース情報を取得
	workspaceQuery := `SELECT id, name, owner_id, theme, created_at, updated_at FROM workspaces WHERE id = $1`
	var workspace Workspace
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

	var channels []Channel
	for rows.Next() {
		var channel Channel
		err := rows.Scan(&channel.ID, &channel.WorkspaceID, &channel.Name, &channel.IsPublic, &channel.CreatedAt, &channel.UpdatedAt)
		if err != nil {
			return nil, err
		}
		channels = append(channels, channel)
	}

	// ワークスペースとチャンネル情報を結合
	response := WorkspaceWithChannels{
		Workspace: workspace,
		Channels:  channels,
	}

	return &response, nil
}