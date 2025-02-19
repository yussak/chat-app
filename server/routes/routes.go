package routes

import (
	"server/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/todos", func(c echo.Context) error {
		return controllers.ListTodos(c)
	})
	e.POST("/todos", func(c echo.Context) error {
		return controllers.AddTodo(c)
	})
	e.DELETE("/todos/:id", func(c echo.Context) error {
		return controllers.DeleteTodo(c)
	})
}
