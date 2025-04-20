package routes

import (
	"server/controllers"
	"server/ui"

	"github.com/labstack/echo/v4"
)

type Handlers struct {
	WorkspaceController  *ui.WorkspaceController
	MessageController    *ui.MessageController
	NavigationController *ui.NavigationController
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

	e.GET("/messages", h.MessageController.GetMessagesHandler)
	e.POST("/messages", h.MessageController.AddMessageHandler)
	e.DELETE("/messages/:id", h.MessageController.DeleteMessageHandler)

	// todo:handlerに揃える
	e.GET("/workspaces", h.WorkspaceController.ListWorkspacesHandler)
	e.POST("/workspaces", h.WorkspaceController.CreateWorkspaceHandler)
	e.GET("/workspaces/:id", h.WorkspaceController.GetWorkspace)

	e.GET("/channels/:id", func(c echo.Context) error {
		return controllers.GetChannel(c)
	})

	// todo:handlerに揃える
	e.GET("/sidebar", h.NavigationController.GetSidebarProps)
}
