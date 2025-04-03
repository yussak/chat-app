package models

import (
	"server/db"
	"time"
)

type Message struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	ChannelID int    `json:"channel_id"`
	User      User   `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	Reactions string `json:"reactions"`
}

// テスト時に関数をモック差し替えできるよう、func ではなく var で定義している
var GetMessages = func(channelID string) ([]Message, error) {
	query := `
		SELECT
			m.id,
			m.content,
			m.created_at,
			m.channel_id,
			u.id,
			u.name,
			u.image,
			COALESCE(
				jsonb_object_agg(r.emoji, r.count) FILTER (WHERE r.emoji IS NOT NULL),
				'{}'::jsonb
			) as reactions
		FROM messages m
		LEFT JOIN users u ON m.user_id = u.id
		LEFT JOIN (
			SELECT message_id, emoji, COUNT(*) as count
			FROM reactions
			GROUP BY message_id, emoji
		) r ON m.id = r.message_id
		WHERE m.channel_id = $1
		GROUP BY m.id, m.content, m.created_at, m.channel_id, u.id, u.name, u.image
		ORDER BY m.created_at ASC
	`
	rows, err := db.DB.Query(query, channelID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message

	for rows.Next() {
		message := Message{}
		user := User{}
		var reactionsJson []byte
		if err := rows.Scan(
			&message.ID,
			&message.Content,
			&message.CreatedAt,
			&message.ChannelID,
			&user.ID,
			&user.Name,
			&user.Image,
			&reactionsJson,
		); err != nil {
			return nil, err
		}
		message.User = user
		message.Reactions = string(reactionsJson)
		messages = append(messages, message)
	}

	return messages, nil
}