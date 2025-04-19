package infrastructure

import (
	"server/db"
	"server/domain"
	"server/models"
	"time"
)

type MessageRepository struct{}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{}
}

type MessageRepositoryImpl struct {}

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
		// todo:domainにする(していいのか調べる)
		user := models.User{}
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

func (r *MessageRepositoryImpl) AddMessage(content string, channelID int, user models.User) (domain.Message, error) {
 // MessagesテーブルにINSERTして、INSERTしたレコードのIDを取得
 var insertedID int
 var createdAt time.Time
 err := db.DB.QueryRow(`INSERT INTO messages (content, user_id, channel_id) VALUES ($1, $2, $3) RETURNING id, created_at`, content, user.ID, channelID).Scan(&insertedID, &createdAt)
 if err != nil {
	 return domain.Message{}, err
 }
 
 // 登録したMessageをJSONで返す
 newMessage := domain.Message{
	 ID:   insertedID,
	 Content: content,
	 User: user,
	 ChannelID: channelID,
	 Reactions: "{}",
	 CreatedAt: createdAt,
 }
 
 return newMessage, nil
}
