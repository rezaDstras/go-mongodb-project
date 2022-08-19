package controller

import (
	"golang-api/model/user"
	"golang-api/service"
	"golang-api/utility"
	userViewModel "golang-api/viewModel/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserById(e echo.Context) error {
	userId := e.Param("id")
	userService := service.NewUserService()
	user, err := userService.GetUserById(userId)
	if err != nil {
		println(err)
	}
	return e.JSON(http.StatusOK, user)
}

func SearchUser(c echo.Context) error {
	userInput := new(user.User)
	err := c.Bind(userInput)
	if err != nil {
		return err
	}
	return nil
	//return c.String(http.StatusOK, "name : "+userInput.Name+" family : "+userInput.Family+" phone : "+userInput.Phone)
}

func Create(e echo.Context) error {
	userInput := new(userViewModel.CreateUserViewModel)
	if err := e.Bind(userInput); err != nil {
		return e.JSON(http.StatusBadRequest, "bad request!!")
	}

	//valodator
	if err:= e.Validate(userInput);err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	//get creator user id => user is defined in jwt in header
	//token := e.Get("user").(*jwt.Token)
	//claim := token.Claims.(*security.JwtClaims)
	//userInput.CreateUserId = claim.ID



	//access to custom context

	apiContext := e.(*utility.ApiContext)
	creatorID , err := apiContext.GetUserId()
	if err != nil {
		return e.JSON(http.StatusForbidden, "unautorized")
	}
	userInput.CreateUserId = creatorID

	userService := service.NewUserService()
	userId, err := userService.InsertUser(*userInput)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "bad request!!")
	}

	userResData := struct {
		NewUserId string
	}{
		NewUserId: userId,
	}
	return e.JSON(http.StatusOK, userResData)
}

func GetAllUsers(e echo.Context) error {
	userService := service.NewUserService()
	userList, err := userService.GetUserList()
	if err != nil {
		println(err)
	}

	return e.JSON(http.StatusOK, userList)
}
