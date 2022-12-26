package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"golang-api-server/api/common"
	"golang-api-server/conf"
	"golang-api-server/db"
	"golang-api-server/route"
)

func settingMiddleware(e *echo.Echo, config conf.Config) {
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}\n",
	// }))

	logImpl := common.TbCoLogSvcImpl{}
	logImpl.Init("api-server", config.Kafka.BootstrapServers)

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		BeforeNextFunc: func(c echo.Context) {
			logImpl.InsertLog("test data")
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			fmt.Printf("REQUEST: uri: %v, status: %v, method: %v\n", v.URI, v.Status, v.Method)
			return nil
		},
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
	settingMiddleware(e, config)

	// routing 설정
	route.InitRoutes(e)

	e.Logger.Fatal(e.Start(":7777"))
}
