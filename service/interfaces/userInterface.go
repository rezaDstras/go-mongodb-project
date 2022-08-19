package interfaces

import (
	user "golang-api/model/user"
	userViewModel "golang-api/viewModel/user"
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
