package routes

import (
	"server/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.POST("/users/signin", func(c echo.Context) error {
		return controllers.SignInHandler(c)
	})

	e.GET("/users/exists", func(c echo.Context) error {
		return controllers.EmailExistsHandler(c)
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

	e.GET("/messages/:id/reactions", func(c echo.Context) error {
		return controllers.ListReactions(c)
	})
	e.POST("/messages/:id/reactions", func(c echo.Context) error {
		return controllers.AddReaction(c)
	})

	e.POST("/workspaces", func(c echo.Context) error {
		return controllers.CreateWorkspace(c)
	})

	e.GET("/workspaces", func(c echo.Context) error {
		return controllers.ListWorkspaces(c)
	})

	e.GET("/workspaces/:id", func(c echo.Context) error {
		return controllers.GetWorkspace(c)
	})
}
