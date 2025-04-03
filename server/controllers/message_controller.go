package controllers

import (
	"net/http"
	"server/db"
	"server/models"

	"github.com/labstack/echo/v4"
)

func ListMessages(c echo.Context) error {
	channelID := c.QueryParam("channel_id")
	if channelID == "" {
		return c.String(http.StatusBadRequest, "ChannelIDが必要です")
	}

	messages, err := models.GetMessages(channelID)
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
	if req.User.ID == 0 {
		return c.String(http.StatusBadRequest, "UserIDが必要です")
	}

	// todo 引数改善
	newMessage, err := models.AddMessage(req.Content, req.ChannelID, req.User)
	if err != nil {
		return c.String(http.StatusInternalServerError, "データベースエラー: " + err.Error())
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
