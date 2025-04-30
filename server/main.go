package main

import (
	"fmt"
	"net/http"
	"server/application"
	"server/db"
	"server/infrastructure"
	"server/routes"
	"server/ui"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

func main() {
	db.Init()

	// todo: 関数に切り出す
	workspaceRepo := infrastructure.NewWorkspaceRepository()
	workspaceService := application.NewWorkspaceService(workspaceRepo)
	workspaceHandler := ui.NewWorkspaceController(workspaceService)

	reactionRepo := infrastructure.NewReactionRepository()
	messageRepo := infrastructure.NewMessageRepository()
	messageService := application.NewMessageService(messageRepo, reactionRepo)
	messageHandler := ui.NewMessageController(messageService)

	navigationRepo := infrastructure.NewNavigationRepository()
	navigationService := application.NewNavigationService(navigationRepo)
	navigationHandler := ui.NewNavigationController(navigationService)

	channelRepo := infrastructure.NewChannelRepository()
	channelService := application.NewChannelService(channelRepo)
	channelHandler := ui.NewChannelController(channelService)

	// ポインタとして保持
	handlers := &routes.Handlers{
		WorkspaceController:  workspaceHandler,
		MessageController:    messageHandler,
		NavigationController: navigationHandler,
		ChannelController:    channelHandler,
	}

	e := echo.New()

	// CORSの設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// アクセスを許可するオリジンを指定
		AllowOrigins: []string{"http://localhost:3000"},
		// アクセスを許可するメソッドを指定
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		// アクセスを許可するヘッダーを指定
		AllowHeaders:     []string{echo.HeaderContentType, echo.HeaderAuthorization, "X-CSRF-Header"},
		AllowCredentials: true,
	}))

	routes.SetupRoutes(e, handlers)

	fmt.Println("Server running on port :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
