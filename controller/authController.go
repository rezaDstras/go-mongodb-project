package controller

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang-api/service"
	"golang-api/viewModel/comman/security"
	"golang-api/viewModel/user"
	"log"
	"net/http"
	"time"
)

func Login(e echo.Context) error   {
	loginModel := new(user.LoginViewModel)

	//bind data
	if err :=e.Bind(loginModel);err !=nil {
		return e.JSON(http.StatusBadRequest,"")
	}

	//validate data
	if err := e.Validate(loginModel);err!=nil{
		return e.JSON(http.StatusBadRequest,"model not valid")
	}

	//TODO Get User

	userService := service.NewUserService()
	user , err := userService.GetUserByUsernameAndPassword(*loginModel)
	if err != nil {
		return e.JSON(http.StatusBadRequest,"user not found")
	}


	//impement jwt
	//create claims from payload
	claims := &security.JwtClaims{
		user.UserName,
		user.Id,
		//we can add eveything we want in claims lik name email or ....
		jwt.StandardClaims{
			//define expire time for jwt(json web token)
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	log.Println(claims.UserName)

	//create token => firt arguman is alguritm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	//encode token => arguman is key for hash token
	tokenString , err := token.SignedString([]byte("secret"))
	if err != nil {
		return e.JSON(http.StatusBadRequest,"token is not valid")
	}

//create response data
	responseData := struct {
		Token string
	}{
		Token: tokenString,

	}
	return e.JSON(http.StatusOK,responseData)
}
