package user

type CreateUserViewModel struct {
	FirstName    string `validate:"required"`
	LastFamily   string
	Email        string
	UserName     string
	Password     string
	CreateUserId string
}
