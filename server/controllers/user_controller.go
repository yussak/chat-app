package controllers

import (
	"log"
	"net/http"
	"server/db"
	"server/models"

	"github.com/labstack/echo/v4"
)

func SignInHandler(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		log.Printf("バインドエラー: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "無効なリクエスト",
		})
	}

	existingUser, err := models.FindUserByEmail(db.DB, user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "ユーザー検索エラー",
		})
	}

	if existingUser == nil {
		// 新規ユーザー作成
		if err := models.CreateUser(db.DB, &user); err != nil {
			log.Printf("ユーザー作成エラー: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "ユーザー作成失敗",
			})
		}
		log.Printf("新規ユーザー作成完了: ID=%d", user.ID)
		existingUser = &user
	} else {
		// 既存ユーザーの更新
		if err := models.UpdateUser(db.DB, &user); err != nil {
			log.Printf("ユーザー更新エラー: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "ユーザー更新失敗",
			})
		}
		log.Println("ユーザー情報更新完了")
		existingUser = &user
	}

	return c.JSON(http.StatusOK, existingUser)
}

func EmailExistsHandler(c echo.Context) error {
	email := c.QueryParam("email")
	if email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "メールアドレスが必要です",
		})
	}

	user, err := models.FindUserByEmail(db.DB, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "ユーザー検索エラー",
		})
	}

	if user == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "ユーザーが見つかりません",
		})
	}

	return c.JSON(http.StatusOK, map[string]bool{
		"exists": true,
	})
}
