package controller

import (
	"golang-api/model/user"
	"golang-api/service"
	"golang-api/utility"
	"golang-api/viewModel/comman/httpResponse"
	userViewModel "golang-api/viewModel/user"
	"io"
	"net/http"
	"os"
	"path/filepath"

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

func Upload(e echo.Context) error {
	apiContext := e.(*utility.ApiContext)

	//get file parameter
	file , err := apiContext.FormFile("file")
	if err != nil {
		return e.JSON(http.StatusBadRequest,httpResponse.ErrorResponse("bad request"))
	}

	//get file
	src ,err := file.Open()
	if err != nil {
		return e.JSON(http.StatusBadRequest,httpResponse.ErrorResponse("can not open file"))
	}

	//create file
	//get working directory
	wd , err := os.Getwd()
	//create custom name for file
	fileName:= "customName"+filepath.Ext(file.Filename) //Ext : get format of file , .png ,.jpg ,...
	imageServerPath := filepath.Join(wd,"public","images",fileName)
	des , err := os.Create(imageServerPath)
	if err != nil {
		return e.JSON(http.StatusBadRequest,httpResponse.ErrorResponse("can not create file"))
	}

	//copy src to des
	_ , err = io.Copy(des,src)
	if err != nil {
		return e.JSON(http.StatusBadRequest,httpResponse.ErrorResponse("can not copy file"))
	}

	userResData := struct {
		IsSuccess bool
	}{
		IsSuccess: true,
	}

	return e.JSON(http.StatusOK,httpResponse.SuccessResponse(userResData))
}
