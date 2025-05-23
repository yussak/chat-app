package routes

import (
	"server/ui"

	"github.com/labstack/echo/v4"
)

type Handlers struct {
	WorkspaceController  *ui.WorkspaceController
	MessageController    *ui.MessageController
	ChannelController    *ui.ChannelController
	NavigationController *ui.NavigationController
	ReactionController   *ui.ReactionController
	UserController       *ui.UserController
}

func SetupRoutes(e *echo.Echo, h *Handlers) {
	e.POST("/users/signin", h.UserController.SignInHandler)

	e.GET("/users/exists", h.UserController.EmailExistsHandler)

	e.GET("/messages/:id/reactions", h.ReactionController.ListReactionsHandler)
	e.POST("/messages/:id/reactions", h.ReactionController.AddReactionHandler)

	e.GET("/messages", h.MessageController.GetMessagesHandler)
	e.POST("/messages", h.MessageController.AddMessageHandler)
	e.DELETE("/messages/:id", h.MessageController.DeleteMessageHandler)

	e.GET("/workspaces", h.WorkspaceController.ListWorkspacesHandler)
	e.POST("/workspaces", h.WorkspaceController.CreateWorkspaceHandler)
	e.GET("/workspaces/:id", h.WorkspaceController.GetWorkspaceHandler)

	e.GET("/channels", h.ChannelController.ListChannelsHandler)

	e.GET("/sidebar", h.NavigationController.GetSidebarPropsHandler)
}
