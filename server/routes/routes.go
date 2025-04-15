package routes

import (
	"server/application"
	"server/controllers"
	"server/infrastructure"
	"server/ui"

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

	workspaceHandler := ui.NewWorkspaceController(application.NewWorkspaceService(infrastructure.NewWorkspaceRepositoryImpl()))
	e.GET("/workspaces", workspaceHandler.ListWorkspaces)
	e.GET("/workspaces/:id", workspaceHandler.GetWorkspace)
	
	e.GET("/channels/:id", func(c echo.Context) error {
		return controllers.GetChannel(c)
	})
	
	navigationHandler := ui.NewNavigationController(application.NewNavigationService(infrastructure.NewNavigationRepositoryImpl()))
	e.GET("/sidebar", navigationHandler.GetSidebarProps)
}
