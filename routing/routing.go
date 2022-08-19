package routing

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"golang-api/routing/routes"
)

func SetRoutes(e *echo.Echo) error {

	//auth routes
	routes.AuthRoutes(e)

	//users routes
	routes.UserRoutes(e)
	//news routes
	routes.NewsRoutes(e)

	//print all routes in consloe
	for i, route := range e.Routes() {
		fmt.Println(i, route.Method, route.Path, route.Name)
	}

	return nil
}
