package ui

import (
	"net/http"
	"server/application"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Service application.UserService
}

func NewUserController(s application.UserService) *UserController {
	return &UserController{Service: s}
}

func (h *UserController) EmailExistsHandler(c echo.Context) error {
	email := c.QueryParam("email")
	if email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "メールアドレスが必要です",
		})
	}

	user, err := h.Service.FindUserByEmail(email)
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
