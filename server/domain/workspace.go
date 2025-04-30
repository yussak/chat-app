package domain

import (
	"database/sql"
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

type WorkspaceOwner struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	Image       string `json:"image"`
	DisplayName string `json:"displayName"`
}

type WorkspaceMember struct {
	ID          int       `json:"id"`
	WorkspaceID int       `json:"workspace_id"`
	UserID      int       `json:"user_id"`
	DisplayName string    `json:"display_name"`
	ImageURL    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type WorkspaceRepository interface {
	FindAll() ([]Workspace, error)
	FindById(id string) (*WorkspaceWithChannels, error)
	// todo: domain依存もダメそうなので直す
	CreateWorkspace(tx *sql.Tx, displayName string, name string, theme string, user *User) error
}
