package ui

import (
	"net/http"
	"server/application"
	"server/db"

	"github.com/labstack/echo/v4"
)

type MessageController struct {
	Service application.MessageService
}

func NewMessageController(s application.MessageService) *MessageController {
	return &MessageController{Service: s}
}

func (h *MessageController) GetMessagesHandler(c echo.Context) error {
	channelID := c.QueryParam("channel_id")
	if channelID == "" {
		return c.String(http.StatusBadRequest, "ChannelIDが必要です")
	}

	messages, err := h.Service.ListMessagesByChannelID(channelID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "データベースエラー: "+err.Error())
	}

	return c.JSON(http.StatusOK, messages)
}

type AddMessageRequest struct {
	Content   string `json:"content"`
	ChannelID int    `json:"channel_id"`
	UserID    int    `json:"user_id"`
}

func (h *MessageController) AddMessageHandler(c echo.Context) error {
	var req AddMessageRequest

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
	if req.UserID == 0 {
		return c.String(http.StatusBadRequest, "UserIDが必要です")
	}

	message, err := h.Service.AddMessage(req.Content, req.ChannelID, req.UserID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "データベースエラー: "+err.Error())
	}

	return c.JSON(http.StatusOK, message)
}

func (h *MessageController) DeleteMessageHandler(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.String(http.StatusBadRequest, "IDが空です")
	}

	// トランザクションを開始
	// todo: db.DB.Begin()はroutes層でやるべきかも uiがdbに依存はダメそう
	tx, err := db.DB.Begin()
	if err != nil {
		return c.String(http.StatusInternalServerError, "トランザクション開始エラー")
	}
	defer tx.Rollback()

	// 関連データを含めて削除
	err = h.Service.DeleteMessageAndRelationData(id, tx)
	if err != nil {
		return c.String(http.StatusInternalServerError, "メッセージ削除エラー")
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return c.String(http.StatusInternalServerError, "トランザクションコミットエラー")
	}

	return c.String(http.StatusOK, "メッセージとリアクションが削除されました")
}
