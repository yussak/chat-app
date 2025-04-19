package models

import (
	"database/sql"
	"server/db"
	"time"
)

type Message struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	ChannelID int       `json:"channel_id"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	Reactions string    `json:"reactions"`
}

func DeleteMessage(id string, tx *sql.Tx) error {
	// まずリアクションを削除
	_, err := db.DB.Exec("DELETE FROM reactions WHERE message_id = $1", id)
	if err != nil {
		return err
	}

	// メッセージを削除
	_, err = tx.Exec("DELETE FROM messages WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
