package domain

import (
	"database/sql"
	"time"
)

type Reaction struct {
	ID        int       `json:"id"`
	MessageID int       `json:"message_id"`
	UserID    int       `json:"user_id"`
	Emoji     string    `json:"emoji"`
	CreatedAt time.Time `json:"created_at"`
}

type ReactionRepository interface {
	Delete(id string, tx *sql.Tx) error
}
