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

func (h *ReactionController) AddReactionHandler(c echo.Context) error {
	messageId := c.Param("id")

	var req struct {
		UserID int    `json:"user_id"`
		Emoji  string `json:"emoji"`
	}

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "リクエストの形式が正しくありません")
	}

	err := h.Service.AddReaction(messageId, req.UserID, req.Emoji)
	if err != nil {
		return c.String(http.StatusInternalServerError, "リアクションの更新に失敗しました: "+err.Error())
	}

	return c.String(http.StatusOK, "リアクションが更新されました")
}
