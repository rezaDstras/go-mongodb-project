package interfaces

import "golang-api/model/user"

type UserRepository interface {
	GetUserList() ([]user.User, error)
	GetUserById(id string) (user.User, error)
	InsertUser(user user.User) (string, error)
	UpdateUser(user user.User) error
	DeleteUserById(id string)error
	GetUserByusernameAndPassword(username , password string) (user.User, error)
}
