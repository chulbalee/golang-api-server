package db

import (
	"database/sql"
	"fmt"
	"golang-api-server/conf"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var once sync.Once
var mysql *sql.DB

func InitDatabase(config conf.Config) {

	once.Do(func() {
		fmt.Println(":::DB CONNECTION SETTING")
		fmt.Println(":::DB Type: ", config.Database.Type)
		fmt.Println(":::DB host: ", config.Database.Host)
		fmt.Println(":::DB NAME: ", config.Database.Name)
		SID := fmt.Sprint(config.Database.User, ":", config.Database.Password, "@tcp(", config.Database.Host, ":", config.Database.Port, ")/", config.Database.Name)
		conn, err := sql.Open(config.Database.Type, SID)

		if err != nil {
			panic(err)
		}

		conn.SetConnMaxIdleTime(time.Minute * 3)
		conn.SetMaxOpenConns(0)
		conn.SetMaxIdleConns(50)

		mysql = conn
	})
}

func GetConn() *sql.DB {
	return mysql
}
