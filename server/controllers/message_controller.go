package controllers

import (
	"net/http"
	"server/db"
	"server/models"

	"github.com/labstack/echo/v4"
)

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
	err = models.DeleteReaction(id, tx)
	if err != nil {
		return c.String(http.StatusInternalServerError, "リアクション削除エラー")
	}

	err = models.DeleteMessage(id, tx)
	if err != nil {
		return c.String(http.StatusInternalServerError, "メッセージ削除エラー")
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return c.String(http.StatusInternalServerError, "トランザクションコミットエラー")
	}

	return c.String(http.StatusOK, "メッセージとリアクションが削除されました")
}
