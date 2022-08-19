package controller

import (
	"fmt"
	"golang-api/service"
	"golang-api/utility"
	"golang-api/viewModel/comman/httpResponse"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"path/filepath"
	newsViewModel "golang-api/viewModel/news"
	uuid "github.com/google/uuid"
)

func GetAllNews(e echo.Context) error {
	newsService := service.NewnewsService()
	newsList, err := newsService.GetNewsList()
	if err != nil {
		fmt.Println(err)
	}

	return e.JSON(http.StatusOK, httpResponse.SuccessResponse(newsList))
}

func CreateNews(e echo.Context) error {
	apiContext := e.(utility.ApiContext)

	newsInput := new(newsViewModel.CreateNewsViewModel)

	//bind data
	if err := apiContext.Bind(newsInput);err != nil {
		return e.JSON(http.StatusBadRequest,httpResponse.ErrorResponse("bad request"))
	}

	//validate
	if err := apiContext.Validate(newsInput);err != nil{
		return e.JSON(http.StatusBadRequest,httpResponse.ErrorResponse("bad request"))
	}

	//upload image

	//get file parameter
	file , err := apiContext.FormFile("file")

	if err == nil {
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

		//generate uniq uuid for image name
		fileName := uuid.New().String()+filepath.Ext(file.Filename) //Ext : get format of file , .png ,.jpg ,...
		imageServerPath := filepath.Join(wd,"public","images",fileName)
		des , err := os.Create(imageServerPath)
		if err != nil {
			return e.JSON(http.StatusBadRequest,httpResponse.ErrorResponse("can not create file"))
		}
		defer des.Close()
		//copy src to des
		_ , err = io.Copy(des,src)
		if err != nil {
			return e.JSON(http.StatusBadRequest,httpResponse.ErrorResponse("can not copy file"))
		}
		newsInput.Image =fileName
	}else{
		return e.JSON(http.StatusBadRequest,httpResponse.ErrorResponse("file not found"))
	}

	newsService := service.NewnewsService()
	news , err := newsService.CreateNews(*newsInput)
	if err != nil {
		return e.JSON(http.StatusBadRequest,httpResponse.ErrorResponse(""))
	}

	newsResDats := struct {
		News string
	}{
		News : news,
	}

	return e.JSON(http.StatusOK,httpResponse.SuccessResponse(newsResDats))
}


