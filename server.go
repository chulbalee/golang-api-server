package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"golang-api-server/conf"
	"golang-api-server/db"
	"golang-api-server/route"
)

func settingMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{""},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
}

func main() {
	var config conf.Config
	var db db.Database

	config = *conf.LoadServerConfig()

	db.InitDatabase(config)
	_ = db

	e := echo.New()

	// 미들웨어 설정()
	settingMiddleware(e)

	// routing 설정
	route.InitRoutes(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":7777"))
}
