package domain

import (
	"server/infrastructure"
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

type WorkspaceRepository interface {
	FindAll() ([]infrastructure.Workspace, error)
}

