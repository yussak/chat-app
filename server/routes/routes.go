package routes

import (
	"server/application"
	"server/controllers"
	"server/infrastructure"
	"server/ui"

	"github.com/labstack/echo/v4"
)

type Handlers struct {
	Workspace *ui.WorkspaceController
	Message   *ui.MessageController
}

func SetupRoutes(e *echo.Echo, h *Handlers) {
	e.POST("/users/signin", func(c echo.Context) error {
		return controllers.SignInHandler(c)
	})

	e.GET("/users/exists", func(c echo.Context) error {
		return controllers.EmailExistsHandler(c)
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

	e.GET("/messages", h.Message.GetMessagesHandler)
	e.POST("/messages", h.Message.AddMessageHandler)
	e.DELETE("/messages/:id", h.Message.DeleteMessageHandler)

	e.GET("/workspaces", h.Workspace.ListWorkspaces)
	e.GET("/workspaces/:id", h.Workspace.GetWorkspace)

	e.GET("/channels/:id", func(c echo.Context) error {
		return controllers.GetChannel(c)
	})

	navigationHandler := ui.NewNavigationController(application.NewNavigationService(infrastructure.NewNavigationRepositoryImpl()))
	e.GET("/sidebar", navigationHandler.GetSidebarProps)
}
