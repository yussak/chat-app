package domain

import (
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
	Reactions string    `json:"reactions"`
}

type MessageRepository interface {
	FindByChannelID(channelID string) ([]Message, error)
	AddMessage(content string, channelID int, userID int) (Message, error)
}
