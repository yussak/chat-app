package routes

import (
	"server/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return controllers.ListTodos(c)
	})
	e.POST("/add", func(c echo.Context) error {
		return controllers.AddTodo(c)
	})
	e.DELETE("/delete", func(c echo.Context) error {
		return controllers.DeleteTodo(c)
	})
}
