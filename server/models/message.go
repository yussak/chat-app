package models

import (
	"database/sql"
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

// テストでモックのためvarとしている
var AddMessage = func(content string, channelID int, user User) (Message, error) {

// MessagesテーブルにINSERTして、INSERTしたレコードのIDを取得
var insertedID int
var createdAt time.Time
err := db.DB.QueryRow(`INSERT INTO messages (content, user_id, channel_id) VALUES ($1, $2, $3) RETURNING id, created_at`, content, user.ID, channelID).Scan(&insertedID, &createdAt)
if err != nil {
	return Message{}, err
}

// 登録したMessageをJSONで返す
newMessage := Message{
	ID:   insertedID,
	Content: content,
	User: user,
	ChannelID: channelID,
	Reactions: "{}",
	CreatedAt: createdAt,
}

return newMessage, nil
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