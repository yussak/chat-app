package controllers

import (
	"net/http"
	"server/db"
	"server/models"
	"time"

	"github.com/labstack/echo/v4"
)

func ListMessages(c echo.Context) error {
	channelID := c.QueryParam("channel_id")
	if channelID == "" {
		return c.String(http.StatusBadRequest, "ChannelIDが必要です")
	}

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
		return c.String(http.StatusInternalServerError, "データベースエラー: " + err.Error())
	}
	defer rows.Close()

	var messages []models.Message

	for rows.Next() {
		message := models.Message{}
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
			return c.String(http.StatusInternalServerError, "データ取得エラー: " + err.Error())
		}
		message.User = user
		message.Reactions = string(reactionsJson)
		messages = append(messages, message)
	}

	return c.JSON(http.StatusOK, messages)
}

func AddMessage(c echo.Context) error {
	var req models.Message

	// JSONボディをバインド
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "リクエストの形式が正しくありません")
	}
	if req.Content == "" {
		return c.String(http.StatusBadRequest, "Messageが空です")
	}
	if req.ChannelID == 0 {
		return c.String(http.StatusBadRequest, "ChannelIDが必要です")
	}

	// MessagesテーブルにINSERTして、INSERTしたレコードのIDを取得
	var insertedID int
	var createdAt time.Time
	err := db.DB.QueryRow(`
		INSERT INTO messages (content, user_id, channel_id) 
		VALUES ($1, $2, $3) 
		RETURNING id, created_at`,
		req.Content,
		req.User.ID,
		req.ChannelID,
	).Scan(&insertedID, &createdAt)

	if err != nil {
		return c.String(http.StatusInternalServerError, "データベースエラー")
	}

	// 登録したMessageをJSONで返す
	newMessage := models.Message{
		ID:   insertedID,
		Content: req.Content,
		User: req.User,
		ChannelID: req.ChannelID,
		Reactions: "{}",
		CreatedAt: createdAt,
	}

	return c.JSON(http.StatusOK, newMessage)
}

func DeleteMessage(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.String(http.StatusBadRequest, "IDが空です")
	}

	// トランザクションを開始
	tx, err := db.DB.Begin()
	if err != nil {
		return c.String(http.StatusInternalServerError, "トランザクション開始エラー")
	}
	defer tx.Rollback()

	// まずリアクションを削除
	_, err = tx.Exec("DELETE FROM reactions WHERE message_id = $1", id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "リアクション削除エラー")
	}

	// メッセージを削除
	_, err = tx.Exec("DELETE FROM messages WHERE id = $1", id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "メッセージ削除エラー")
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return c.String(http.StatusInternalServerError, "トランザクションコミットエラー")
	}

	return c.String(http.StatusOK, "メッセージとリアクションが削除されました")
}
