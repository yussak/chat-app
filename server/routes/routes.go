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

	e.GET("/messages/:id/reactions", func(c echo.Context) error {
		return controllers.ListReactions(c)
	})
	e.POST("/messages/:id/reactions", func(c echo.Context) error {
		return controllers.AddReaction(c)
	})

	e.POST("/workspaces", func(c echo.Context) error {
		return controllers.CreateWorkspace(c)
	})

	// todo:一気に変更せず変数作りながらやりたい messageRepo := ..のように
	workspaceHandler := ui.NewWorkspaceController(application.NewWorkspaceService(infrastructure.NewWorkspaceRepository()))
	messageHandler := ui.NewMessageController(application.NewMessageService(infrastructure.NewMessageRepository(), infrastructure.NewReactionRepository()))

	e.GET("/messages", messageHandler.GetMessagesHandler)
	e.POST("/messages", messageHandler.AddMessageHandler)
	e.DELETE("/messages/:id", messageHandler.DeleteMessageHandler)

	e.GET("/workspaces", workspaceHandler.ListWorkspaces)
	e.GET("/workspaces/:id", workspaceHandler.GetWorkspace)

	e.GET("/channels/:id", func(c echo.Context) error {
		return controllers.GetChannel(c)
	})

	navigationHandler := ui.NewNavigationController(application.NewNavigationService(infrastructure.NewNavigationRepositoryImpl()))
	e.GET("/sidebar", navigationHandler.GetSidebarProps)
}
