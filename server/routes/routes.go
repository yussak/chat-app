package routes

import (
	"server/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.POST("/users/signin", func(c echo.Context) error {
		return controllers.SignInHandler(c)
	})

	e.GET("/messages", func(c echo.Context) error {
		return controllers.ListMessages(c)
	})
	e.POST("/messages", func(c echo.Context) error {
		return controllers.AddMessage(c)
	})
	e.DELETE("/messages/:id", func(c echo.Context) error {
		return controllers.DeleteMessage(c)
	})
}
