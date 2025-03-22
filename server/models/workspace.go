package models

import "time"

type Workspace struct {
	ID int `json:"id"`
	Email string `json:"email"`
	OwnerID int `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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