package ui

import (
	"net/http"
	"server/application"

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