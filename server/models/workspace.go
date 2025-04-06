package models

import (
	"database/sql"
	"fmt"
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

func CreateWorkspace(tx *sql.Tx, workspace *Workspace, displayName string, user *User) error {
	// ワークスペースを作成
	query := `INSERT INTO workspaces (owner_id, name, theme) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
	err := tx.QueryRow(query, workspace.OwnerID, workspace.Name, workspace.Theme).Scan(&workspace.ID, &workspace.CreatedAt, &workspace.UpdatedAt)
	if err != nil {
		return fmt.Errorf("ワークスペース作成エラー: %w", err)
	}

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