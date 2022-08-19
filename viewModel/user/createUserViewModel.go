package user

type CreateUserViewModel struct {
	FirstName    string `validate:"required"`
	LastFamily   string
	Email        string
	UserName     string
	Password     string
	CreateUserId string
}

type UpdateUserViewModel struct {
	TargetUserID string
	FirstName    string `validate:"required"`
	LastFamily   string `validate:"required"`
	Email        string `validate:"required"`
	UserName     string `validate:"required"`
	Password     string `validate:"required"`
}
