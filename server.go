package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"myapp/route"
)

const (
	SERVER_CONFIG_PATH = "./conf/server.json"
)

type Config struct {
	DB_DRIVER string `json:"db.driver"`
	DB_ADDR   string `json:"db.addr"`
	DB_PORT   string `json:"db.port"`
	DB_USER   string `json:"db.user"`
	DB_PWD    string `json:"db.pwd"`
	DB_NAME   string `json:"db.name"`
}

func loadServerConfig() Config {
	var config Config
	fi, err := os.Open(SERVER_CONFIG_PATH)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	jsonParser := json.NewDecoder(fi)
	jsonParser.Decode(&config)

	return config
}

func loadDbConnection(config Config) sql.DB {
	fmt.Println(":::DB CONNECTION SETTING")
	fmt.Println(":::DB DRIVER: ", config.DB_DRIVER)
	fmt.Println(":::DB ADDR: ", config.DB_ADDR)
	fmt.Println(":::DB NAME: ", config.DB_NAME)

	SID := fmt.Sprint(config.DB_USER, ":", config.DB_PWD, "@tcp(", config.DB_ADDR, ":", config.DB_PORT, ")/", config.DB_NAME)
	db, err := sql.Open(config.DB_DRIVER, SID)

	if err != nil {
		panic(err)
	}

	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetMaxOpenConns(0)
	db.SetMaxIdleConns(50)

	return *db
}

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
	e := echo.New()

	// 설정파일 로드
	config := loadServerConfig()

	// db conn 설정
	db := loadDbConnection(config)
	_ = db

	// 미들웨어 설정()
	settingMiddleware(e)

	// routing 설정
	route.InitRoutes(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":7777"))
}
