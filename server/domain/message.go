package domain

import (
	"server/models"
	"time"
)

type Message struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	ChannelID int    `json:"channel_id"`
	// todo:domainにするのはよくなさそうなので自分の所でUserを用意する予定？UserIdはいいけどUserはNGらしい　アプリ層でやるとかなんとか。調べる
	User      models.User   `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	Reactions string `json:"reactions"`
}

type MessageRepository interface {
	FindByChannelID(channelID string) ([]Message, error)
	// todo:modelsの依存をなくす
	AddMessage(content string, channelID int, user models.User) (Message, error)
}