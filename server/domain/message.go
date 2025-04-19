package domain

import (
	"database/sql"
	"time"
)

type UserInfo struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type Message struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	ChannelID int       `json:"channel_id"`
	User      UserInfo  `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	// todo:これstringでいいのか？
	Reactions string `json:"reactions"`
}

// todo:string intなどはよくなさそうなので改善
type MessageRepository interface {
	FindByChannelID(channelID string) ([]Message, error)
	AddMessage(content string, channelID int, userID int) (Message, error)
	Delete(id string, tx *sql.Tx) error
}
