package routing

import (
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"golang-api/controller"
	"golang-api/viewModel/comman/security"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetRoutes(e *echo.Echo) error {

	//regular routing
	e.GET("/user/15", func(c echo.Context) error {
		return c.String(http.StatusOK, "/user/15")
	}).Name = "get-user-15"


	//login
	e.POST("/login",controller.Login)

	//grouping routes
	user := e.Group("/user/")
	user.GET(":id", controller.GetUserById).Name = "get-user-id"
	user.GET("search", controller.SearchUser).Name = "search-user"


	jwtConfig:=middleware.JWTConfig{
		//secret key for jwt which is set in login
		SigningKey: []byte("secret"),
		Claims: &security.JwtClaims{},
	}

	user.POST("create", controller.Create,middleware.JWTWithConfig(jwtConfig)).Name = "create-user"



	user.GET("list", controller.GetAllUsers).Name = "user-list"

	//print all routes in consloe
	for i, route := range e.Routes() {
		fmt.Println(i, route.Method, route.Path, route.Name)
	}

	return nil
}
