package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetArticles(c echo.Context) error {

	data := `{"status":200}`

	return c.JSON(http.StatusOK, data)
}

func CreateArticle(c echo.Context) error {
	data := `{"status":200}`

	return c.JSON(http.StatusOK, data)
}
