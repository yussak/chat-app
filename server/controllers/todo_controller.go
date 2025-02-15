package controllers

import (
	"net/http"
	"server/db"
	"server/models"

	"github.com/labstack/echo/v4"
)

func ListTodos(c echo.Context) error {
	rows, err := db.DB.Query("SELECT id, name FROM todos")
	if err != nil {
		return c.String(http.StatusInternalServerError, "データベースエラー")
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		todo := models.Todo{}
		if err := rows.Scan(&todo.ID, &todo.Name); err != nil {
			return c.String(http.StatusInternalServerError, "データ取得エラー")
		}
		todos = append(todos, todo)
	}

	return c.JSON(http.StatusOK, todos)
}

func AddTodo(c echo.Context) error {
	var req models.Todo

	// JSONボディをバインド
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "リクエストの形式が正しくありません")
	}
	if req.Name == "" {
		return c.String(http.StatusBadRequest, "TODO名が空です")
	}

	// TodosテーブルにINSERTして、INSERTしたレコードのIDを取得
	var insertedID int
	err := db.DB.QueryRow(
		"INSERT INTO Todos (name) VALUES ($1) RETURNING id",
		req.Name,
	).Scan(&insertedID)

	if err != nil {
		return c.String(http.StatusInternalServerError, "データベースエラー")
	}

	// 登録したTODOをJSONで返す
	newTodo := models.Todo{
		ID:   insertedID,
		Name: req.Name,
	}

	return c.JSON(http.StatusOK, newTodo)
}

func DeleteTodo(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return c.String(http.StatusBadRequest, "IDが空です")
	}

	// データベースから削除
	_, err := db.DB.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "データベースエラー")
	}

	return c.Redirect(http.StatusSeeOther, "/")
}
