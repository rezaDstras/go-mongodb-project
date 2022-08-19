package service

import (
	"golang-api/model/user"
	"golang-api/repository"
	userViewModel "golang-api/viewModel/user"
	"time"
)

type UserService interface {
	GetUserList() ([]user.User, error)
	GetUserById(id string) (user.User, error)
	GetUserByUsernameAndPassword(login userViewModel.LoginViewModel) (user.User, error)
	InsertUser(userInput userViewModel.CreateUserViewModel) (string ,error)
	UpadteUser(user user.User) error
	DeleteUserByID(id string) error
}

type userService struct {
}

func NewUserService() UserService {
	return userService{}
}

func (u userService) GetUserList() ([]user.User, error) {
	userRepository := repository.NewUserRepository()
	userList , err := userRepository.GetUserList()
	return userList,err
}
func (u userService)GetUserById(id string)(user.User , error)  {
	userRepoitory := repository.NewUserRepository()
	user , err := userRepoitory.GetUserById(id)
	return user,err
}

func (u userService)InsertUser(userInput userViewModel.CreateUserViewModel) (string ,error)  {
	userEntity := user.User{
		FirstName:    userInput.FirstName,
		LastFamily:   userInput.LastFamily,
		Email:        userInput.Email,
		UserName:     userInput.UserName,
		Password:     userInput.Password,
		CreatorUserId: userInput.CreateUserId,
		RegisterDate: time.Now(),
	}
	userRepoitory := repository.NewUserRepository()
	userId, err := userRepoitory.InsertUser(userEntity)
	return userId,err
}
func (u userService)UpadteUser(user user.User) error  {
	userRepoitory := repository.NewUserRepository()
	 err := userRepoitory.UpdateUser(user)
	return err
}
func (u userService)DeleteUserByID(id string) error  {
	userRepoitory := repository.NewUserRepository()
	err := userRepoitory.DeleteUserById(id)
	return err
}
func (u userService)GetUserByUsernameAndPassword(login userViewModel.LoginViewModel) (user.User , error)  {
	userRepoitory := repository.NewUserRepository()
	user , err := userRepoitory.GetUserByusernameAndPassword(login.UserName,login.Password)
	return user ,err
}
