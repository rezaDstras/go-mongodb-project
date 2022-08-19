package routes

import (
	"github.com/labstack/echo/v4"
	"golang-api/controller"
)

func NewsRoutes(e *echo.Echo)  {
	news := e.Group("/news")
	news.GET("/list",controller.GetAllNews).Name="news-list"
	news.POST("/create",controller.CreateNews).Name="news-create"
}