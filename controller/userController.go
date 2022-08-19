package controller

import (
	"golang-api/model/user"
	"golang-api/service"
	"golang-api/utility"
	"golang-api/viewModel/comman/httpResponse"
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
	return e.JSON(http.StatusOK,httpResponse.SuccessResponse(user))
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

	apiContext := e.(*utility.ApiContext)
	creatorID , err := apiContext.GetUserId()
	if err != nil {
		return e.JSON(http.StatusForbidden, httpResponse.ErrorResponse("unautorized"))
	}
	userService := service.NewUserService()
	//check role access
	isValid := userService.IsUserValidForAccess(creatorID,"createUser")
	if !isValid{
		return e.JSON(http.StatusForbidden, httpResponse.ErrorResponse( "you don not access!!"))
	}

	userInput := new(userViewModel.CreateUserViewModel)
	if err := e.Bind(userInput); err != nil {
		return e.JSON(http.StatusBadRequest, httpResponse.ErrorResponse("bad request!!"))
	}

	//valodator
	if err:= e.Validate(userInput);err != nil {
		return e.JSON(http.StatusBadRequest, httpResponse.ErrorResponse("bad request"))
	}

	userInput.CreateUserId = creatorID

	userId, err := userService.InsertUser(*userInput)
	if err != nil {
		return e.JSON(http.StatusBadRequest, httpResponse.ErrorResponse("bad request!!"))
	}

	userResData := struct {
		NewUserId string
	}{
		NewUserId: userId,
	}
	return e.JSON(http.StatusOK, httpResponse.SuccessResponse(userResData))
}

func Update(e echo.Context) error  {
	apiContext := e.(*utility.ApiContext)
	//get id parameter
	targetUserId := apiContext.Param("id")

	userService := service.NewUserService()

	//get data
	updateUser := new(userViewModel.UpdateUserViewModel)

	//bind data
	if err := e.Bind(updateUser);err != nil {
		return e.JSON(http.StatusBadRequest,httpResponse.ErrorResponse("bad request"))
	}

	//validate data
	if err := e.Validate(updateUser); err != nil {
		return e.JSON(http.StatusBadRequest,httpResponse.ErrorResponse("bad request"))
	}

	updateUser.TargetUserID = targetUserId

	//check user is exist or not
	if !userService.IsExistUser(targetUserId){
		return e.JSON(http.StatusNotFound,httpResponse.NotFoundResponse("user not found"))
	}

	err := userService.UpadteUser(*updateUser)
	if err != nil {
		return e.JSON(http.StatusBadRequest,httpResponse.ErrorResponse("bad request"))
	}

	userResData := struct {
		IsSuccess bool
	}{
		IsSuccess: true,
	}

	return e.JSON(http.StatusOK,httpResponse.SuccessResponse(userResData))
}

func GetAllUsers(e echo.Context) error {
	userService := service.NewUserService()
	userList, err := userService.GetUserList()
	if err != nil {
		println(err)
	}

	return e.JSON(http.StatusOK, httpResponse.SuccessResponse(userList))
}

func Delete(e echo.Context) error {
	apiContext := e.(*utility.ApiContext)

	targetUserID := apiContext.Param("id")

	userService := service.NewUserService()
	//check user is exist or not
	if !userService.IsExistUser(targetUserID){
		return e.JSON(http.StatusNotFound,"user not found")
	}

	if err:= userService.DeleteUserByID(targetUserID);err != nil {
		return e.JSON(http.StatusBadRequest,"")
	}

	userResData := struct {
		IsSuccess bool
	}{
		IsSuccess: true,
	}

	return e.JSON(http.StatusOK,httpResponse.SuccessResponse(userResData))



}
