package controllers

import (
	"net/http"
	"server/db"
	"server/models"

	"github.com/labstack/echo/v4"
)

func ListMessages(c echo.Context) error {
	query := `
		SELECT m.id, m.content, u.id, u.name, u.image 
		FROM messages m 
		LEFT JOIN users u ON m.user_id = u.id
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return c.String(http.StatusInternalServerError, "データベースエラー: " + err.Error())
	}
	defer rows.Close()

	var messages []models.Message

	for rows.Next() {
		message := models.Message{}
		user := models.User{}
		if err := rows.Scan(
			&message.ID, 
			&message.Content, 
			&user.ID, 
			&user.Name, 
			&user.Image,
		); err != nil {
			return c.String(http.StatusInternalServerError, "データ取得エラー: " + err.Error())
		}
		message.User = user
		messages = append(messages, message)
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

	// MessagesテーブルにINSERTして、INSERTしたレコードのIDを取得
	var insertedID int
	err := db.DB.QueryRow(
		"INSERT INTO messages (content, user_id) VALUES ($1, $2) RETURNING id",
		req.Content,
		req.User.ID,
	).Scan(&insertedID)

	if err != nil {
		return c.String(http.StatusInternalServerError, "データベースエラー")
	}

	// 登録したMessageをJSONで返す
	newMessage := models.Message{
		ID:   insertedID,
		Content: req.Content,
		User: req.User,
	}

	return c.JSON(http.StatusOK, newMessage)
}

func DeleteMessage(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.String(http.StatusBadRequest, "IDが空です")
	}

	// データベースから削除
	_, err := db.DB.Exec("DELETE FROM messages WHERE id = $1", id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "データベースエラー")
	}

	return c.String(http.StatusOK, "Messageが削除されました")
}
