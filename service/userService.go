package service

import (
	"golang-api/model/user"
	"golang-api/repository"
	userViewModel "golang-api/viewModel/user"
	"time"
	"golang.org/x/exp/slices"
)

type UserService interface {
	GetUserList() ([]user.User, error)
	GetUserById(id string) (user.User, error)
	GetUserByUsernameAndPassword(login userViewModel.LoginViewModel) (user.User, error)
	InsertUser(userInput userViewModel.CreateUserViewModel) (string ,error)
	UpadteUser(userInput userViewModel.UpdateUserViewModel) error
	DeleteUserByID(id string) error
	IsUserValidForAccess(userId , roleName string) bool
	IsExistUser(userId string) bool
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
func (u userService)UpadteUser(userInput userViewModel.UpdateUserViewModel) error {

	userEntity := user.User{
		Id: userInput.TargetUserID,
		FirstName:    userInput.FirstName,
		LastFamily:   userInput.LastFamily,
		Email:        userInput.Email,
		UserName:     userInput.UserName,
		Password:     userInput.Password,
	}


	userRepoitory := repository.NewUserRepository()
	 err := userRepoitory.UpdateUser(userEntity)
	return err
}
func (u userService)DeleteUserByID(id string) error  {
	userRepoitory := repository.NewUserRepository()
	err := userRepoitory.DeleteUserById(id)
	return err
}
func (u userService)GetUserByUsernameAndPassword(login userViewModel.LoginViewModel) (user.User , error) {
	userRepoitory := repository.NewUserRepository()
	user, err := userRepoitory.GetUserByusernameAndPassword(login.UserName, login.Password)
	return user, err
}
func (u userService) IsUserValidForAccess(userId , roleName string) bool{
	userRepository := repository.NewUserRepository()
	user , err := userRepository.GetUserById(userId)
	if err != nil {
		return false
	}

	if user.Roles == nil{
		return false
	}

	//search in array with slices package

	//if value exist in array return value if not return -1
	roleIndex := slices.IndexFunc(user.Roles, func(role string) bool {
		return role == roleName
	})
	 return roleIndex >= 0
}

func (u userService) IsExistUser(userId string) bool {
	userRepository := repository.NewUserRepository()

	_, err := userRepository.GetUserById(userId)
	if err != nil {
		return false
	}

	return true
}
