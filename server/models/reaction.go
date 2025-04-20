package models

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

func DeleteReaction(id string, tx *sql.Tx) error {
	_, err := tx.Exec("DELETE FROM reactions WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
