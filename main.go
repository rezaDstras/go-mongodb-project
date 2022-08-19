package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4/middleware"
	"golang-api/config"
	"golang-api/routing"
	"golang-api/utility"
	"golang-api/viewModel/comman/security"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	// get cinfig
	err := config.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Server Port :", config.AppConfig.Server.Port)

	// initial server
	server := echo.New()
	//validate request use custom validator
	server.Validator = &utility.CustomValidator{Validator : validator.New()}

	//routing
	routing.SetRoutes(server)

	//middleware

	//customize context
	server.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			apiContext := &utility.ApiContext{Context:c}
			return next(apiContext)
		}
	})

	// check jwt
	server.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		//secret key for jwt which is set in login
		SigningKey: []byte("secret"),
		Claims: &security.JwtClaims{},
		//if user not utorize continue
		ContinueOnIgnoredError: true,
		ErrorHandlerWithContext: func(err error, c echo.Context) error {
			//return nil ignore error to continue
			return nil
		},
	}))

	//trottle => each user can 20 request every second
	server.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))



	//start server
	err = server.Start(":" + config.AppConfig.Server.Port)
	if err != nil {
		return
	}

}
