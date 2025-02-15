package main

import (
	"fmt"
	"net/http"
	"server/db"
	"server/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

func main() {
	db.Init()

	e := echo.New()

	// CORSの設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// アクセスを許可するオリジンを指定
		AllowOrigins: []string{"http://localhost:5173"},
		// アクセスを許可するメソッドを指定
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		// アクセスを許可するヘッダーを指定
		AllowHeaders:     []string{echo.HeaderContentType, echo.HeaderAuthorization, "X-CSRF-Header"},
		AllowCredentials: true,
	}))

	routes.SetupRoutes(e)

	fmt.Println("Server running on port :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
