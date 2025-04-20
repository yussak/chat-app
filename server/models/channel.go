package models

import (
	"server/db"
	"time"
)

type Channel struct {
	ID          int64     `json:"id"`
	WorkspaceID int64     `json:"workspace_id"`
	Name        string    `json:"name"`
	IsPublic    bool      `json:"is_public"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func GetChannel(id string) (Channel, error) {
	query := `SELECT id, name, created_at, updated_at FROM channels WHERE id = $1`

	var channel Channel
	err := db.DB.QueryRow(query, id).Scan(
		&channel.ID,
		&channel.Name,
		&channel.CreatedAt,
		&channel.UpdatedAt,
	)
	if err != nil {
		return Channel{}, err
	}

	return channel, nil
}
