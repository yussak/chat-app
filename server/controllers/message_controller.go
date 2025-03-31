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

	messages, err := models.ListMessages(channelID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "データベースエラー: " + err.Error())
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
