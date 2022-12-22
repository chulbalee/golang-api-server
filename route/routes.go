package route

import (
	"myapp/handler"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {

	api := e.Group("/api/v1", serverHeader)
	api.GET("/articles", handler.GetArticles)
	api.POST("/articles", handler.CreateArticle)
}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("ver", "v1")
		return next(c)
	}
}
