package routes

import (
	"github.com/labstack/echo/v4"
	"golang-api/controller"
)

func AuthRoutes(e * echo.Echo)  {
	e.POST("/login",controller.Login)
}
