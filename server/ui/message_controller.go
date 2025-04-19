package ui

import (
	"net/http"
	"server/application"
	"server/models"

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
		return c.String(http.StatusInternalServerError, "データベースエラー: " + err.Error())
	}

	return c.JSON(http.StatusOK, messages)
}

func (h *MessageController) AddMessageHandler(c echo.Context) error {
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

	newMessage, err := models.AddMessage(req.Content, req.ChannelID, req.User)
	if err != nil {
		return c.String(http.StatusInternalServerError, "データベースエラー: " + err.Error())
	}

	return c.JSON(http.StatusOK, newMessage)
}