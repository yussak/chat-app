package domain

import (
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

type WorkspaceSidebarProps struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	YoungestChannelID int64  `json:"youngestChannelId"`
}

type WorkspaceRepository interface {
	FindAll() ([]Workspace, error)
	FindById(id string) (*WorkspaceWithChannels, error)
	GetSidebarProps() ([]WorkspaceSidebarProps, error)
}
