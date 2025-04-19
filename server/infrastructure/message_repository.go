package infrastructure

import (
	"database/sql"
	"server/db"
	"server/domain"
	"time"
)

type MessageRepository struct{}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{}
}

type MessageRepositoryImpl struct{}

func NewMessageRepositoryImpl() domain.MessageRepository {
	return &MessageRepositoryImpl{}
}

func (r *MessageRepositoryImpl) FindByChannelID(channelID string) ([]domain.Message, error) {
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

	var messages []domain.Message

	for rows.Next() {
		message := domain.Message{}
		user := domain.UserInfo{}
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

func (r *MessageRepositoryImpl) AddMessage(content string, channelID int, userID int) (domain.Message, error) {
	var id int
	var createdAt time.Time

	err := db.DB.QueryRow(`INSERT INTO messages (content, user_id, channel_id) VALUES ($1, $2, $3) RETURNING id, created_at`, content, userID, channelID).Scan(&id, &createdAt)
	if err != nil {
		return domain.Message{}, err
	}

	var user domain.UserInfo
	err = db.DB.QueryRow(`SELECT id, name, image FROM users WHERE id = $1`, userID).Scan(&user.ID, &user.Name, &user.Image)
	if err != nil {
		return domain.Message{}, err
	}

	newMessage := domain.Message{
		ID:        id,
		Content:   content,
		User:      user,
		ChannelID: channelID,
		Reactions: "{}",
		CreatedAt: createdAt,
	}

	return newMessage, nil
}

func (r *MessageRepositoryImpl) Delete(id string, tx *sql.Tx) error {
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
