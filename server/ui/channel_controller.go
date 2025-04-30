package ui

import (
	"net/http"
	"server/application"

	"github.com/labstack/echo/v4"
)

type ChannelController struct {
	Service application.ChannelService
}

func NewChannelController(s application.ChannelService) *ChannelController {
	return &ChannelController{Service: s}
}

func (h *ChannelController) ListChannelsHandler(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID is required"})
	}

	channel, err := h.Service.GetChannel(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "channel not found"})
	}

	return c.JSON(http.StatusOK, channel)
}
