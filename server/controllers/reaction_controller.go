package controllers

import (
	"net/http"
	"server/db"

	"github.com/labstack/echo/v4"
)

func ListReactions(c echo.Context) error {
	messageId := c.Param("id")
	rows, err := db.DB.Query(`
	SELECT emoji, COUNT(*)
	FROM reactions
	WHERE message_id = $1
	GROUP BY emoji
	`, messageId)

	if err != nil {
		return c.String(http.StatusInternalServerError, "データベースエラー")
	}

	defer rows.Close()

	type Reaction struct {
		Emoji string `json:"emoji"`
		Count int    `json:"count"`
	}

	reactions := []Reaction{}
	for rows.Next() {
		var reaction Reaction
		if err := rows.Scan(&reaction.Emoji, &reaction.Count); err != nil {
			return c.String(http.StatusInternalServerError, "データ取得エラー")
		}
		reactions = append(reactions, reaction)
	}

	return c.JSON(http.StatusOK, reactions)
}

func AddReaction(c echo.Context) error {
	messageId := c.Param("id")

	var req struct {
		UserID int    `json:"user_id"`
		Emoji  string `json:"emoji"`
	}

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "リクエストの形式が正しくありません")
	}

	// 既存のリアクションを確認
	var exists bool
	err := db.DB.QueryRow(`
		SELECT EXISTS (
			SELECT 1 FROM reactions 
			WHERE message_id = $1 AND user_id = $2 AND emoji = $3
		)`,
		messageId, req.UserID, req.Emoji,
	).Scan(&exists)

	if err != nil {
		return c.String(http.StatusInternalServerError, "データベースエラー: "+err.Error())
	}

	if exists {
		// リアクションが存在する場合は削除
		_, err = db.DB.Exec(`
			DELETE FROM reactions 
			WHERE message_id = $1 AND user_id = $2 AND emoji = $3`,
			messageId, req.UserID, req.Emoji,
		)
	} else {
		// リアクションが存在しない場合は追加
		_, err = db.DB.Exec(`
			INSERT INTO reactions (message_id, user_id, emoji)
			VALUES ($1, $2, $3)`,
			messageId, req.UserID, req.Emoji,
		)
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "リアクションの更新に失敗しました: "+err.Error())
	}

	return c.String(http.StatusOK, "リアクションが更新されました")
}
