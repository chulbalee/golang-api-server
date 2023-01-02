package handler

import (
	"golang-api-server/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetArticles(c echo.Context) error {

	rows, _ := db.GetConn().Query("SELECT * FROM TB_CO_API_HIST")

	columns, _ := rows.Columns()

	return c.JSON(http.StatusOK, columns)
}

func CreateArticle(c echo.Context) error {
	data := `{"status":200}`

	return c.JSON(http.StatusOK, data)
}
