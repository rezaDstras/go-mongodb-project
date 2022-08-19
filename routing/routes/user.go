package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang-api/controller"
	"golang-api/viewModel/comman/security"
)

func UserRoutes(e *echo.Echo)  {

	jwtConfig:=middleware.JWTConfig{
		//secret key for jwt which is set in login
		SigningKey: []byte("secret"),
		Claims: &security.JwtClaims{},
	}

	user := e.Group("/user/")
	user.GET(":id", controller.GetUserById).Name = "get-user-id"
	user.GET("search", controller.SearchUser).Name = "search-user"

	user.POST("create", controller.Create,middleware.JWTWithConfig(jwtConfig)).Name = "create-user"

	user.GET("list", controller.GetAllUsers).Name = "user-list"
	user.PUT("update/:id", controller.Update).Name = "user-update"
	user.DELETE("delete/:id", controller.Delete).Name = "user-delete"
	user.POST("upload", controller.Upload).Name = "user-upload"
}
