package ui

import (
	"net/http"
	"server/application"

	"github.com/labstack/echo/v4"
)

type ReactionController struct {
	Service application.ReactionService
}

func NewReactionController(s application.ReactionService) *ReactionController {
	return &ReactionController{Service: s}
}

func (h *ReactionController) ListReactionsHandler(c echo.Context) error {
	messageId := c.Param("id")
	reactions, err := h.Service.ListReactions(messageId)
	if err != nil {
		return c.String(http.StatusInternalServerError, "データベースエラー")
	}

	return c.JSON(http.StatusOK, reactions)
}
