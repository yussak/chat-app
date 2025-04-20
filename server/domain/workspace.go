package domain

import (
	"database/sql"
	"server/models"
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

type WorkspaceRepository interface {
	FindAll() ([]Workspace, error)
	FindById(id string) (*WorkspaceWithChannels, error)
	// todo:models依存を廃止
	// CreateWorkspace(tx *sql.Tx, displayName string, name string, theme string, user *WorkspaceOwner) error
	CreateWorkspace(tx *sql.Tx, displayName string, name string, theme string, user *models.User) error
}
