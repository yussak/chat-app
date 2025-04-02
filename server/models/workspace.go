package models

import (
	"database/sql"
	"fmt"
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

type WorkspaceMember struct {
	ID int `json:"id"`
	WorkspaceID int `json:"workspace_id"`
	UserID int `json:"user_id"`
	DisplayName string `json:"display_name"`
	ImageURL string `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ListWorkspaces() ([]Workspace, error) {
query := `
		SELECT
			w.id,
			w.name,
			w.owner_id,
			w.theme,
			w.created_at,
			w.updated_at
		FROM workspaces w
	`
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

func GetWorkspace(id string) (*WorkspaceWithChannels, error) {
	// ワークスペース情報を取得
	workspaceQuery := `
		SELECT 
			id,
			name,
			owner_id,
			theme,
			created_at,
			updated_at
		FROM workspaces
		WHERE id = $1
	`
	var workspace Workspace
	err := db.DB.QueryRow(workspaceQuery, id).Scan(
		&workspace.ID,
		&workspace.Name,
		&workspace.OwnerID,
		&workspace.Theme,
		&workspace.CreatedAt,
		&workspace.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	// チャンネル情報を取得
	channelsQuery := `
		SELECT 
			id,
			workspace_id,
			name,
			is_public,
			created_at,
			updated_at
		FROM channels
		WHERE workspace_id = $1
	`
	rows, err := db.DB.Query(channelsQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var channels []Channel
	for rows.Next() {
		var channel Channel
		err := rows.Scan(
			&channel.ID,
			&channel.WorkspaceID,
			&channel.Name,
			&channel.IsPublic,
			&channel.CreatedAt,
			&channel.UpdatedAt,
		)
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

func CreateWorkspaceWithChannels(tx *sql.Tx, workspace *Workspace, displayName string, user *User) error {
	// ワークスペースを作成
	query := `INSERT INTO workspaces (owner_id, name, theme) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
	err := tx.QueryRow(query, workspace.OwnerID, workspace.Name, workspace.Theme).Scan(&workspace.ID, &workspace.CreatedAt, &workspace.UpdatedAt)
	if err != nil {
		return fmt.Errorf("ワークスペース作成エラー: %w", err)
	}

	// // ワークスペースメンバーを作成
	err = CreateWorkspaceMember(tx, workspace, user, displayName)
	if err != nil {
		return fmt.Errorf("ワークスペースメンバー作成エラー: %w", err)
	}

	err = CreateDefaultChannel(tx, workspace)
	if err != nil {
		return fmt.Errorf("チャンネル作成エラー: %w", err)
	}

	return nil
}

func CreateWorkspaceMember(tx *sql.Tx, workspace *Workspace, user *User, displayName string) error {
		workspaceMember := WorkspaceMember{
			WorkspaceID: workspace.ID,
			UserID:      user.ID,
			DisplayName: displayName,
			ImageURL:    user.Image,
		}

		query := `INSERT INTO workspace_members (workspace_id, user_id, display_name, image_url) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`
		err := tx.QueryRow(query, workspaceMember.WorkspaceID, workspaceMember.UserID, workspaceMember.DisplayName, workspaceMember.ImageURL).Scan(&workspaceMember.ID, &workspaceMember.CreatedAt, &workspaceMember.UpdatedAt)
		if err != nil {
			return err
		}

		return nil
}

func CreateDefaultChannel(tx *sql.Tx, workspace *Workspace) error {
	channels := []Channel{
		{
			WorkspaceID: int64(workspace.ID),
			Name:        fmt.Sprintf("all-%s", workspace.Name),
			IsPublic:    true,
		},
		{
			WorkspaceID: int64(workspace.ID),
			Name:        "ソーシャル",
			IsPublic:    true,
		},
		{
			WorkspaceID: int64(workspace.ID),
			Name:        workspace.Theme,
			IsPublic:    true,
		},
	}

	for _, channel := range channels {
		query := `INSERT INTO channels (workspace_id, name, is_public) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
		err := tx.QueryRow(query, channel.WorkspaceID, channel.Name, channel.IsPublic).Scan(&channel.ID, &channel.CreatedAt, &channel.UpdatedAt)
		if err != nil {
			return err
		}
	}

	return nil
}