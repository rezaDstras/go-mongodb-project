package news

type CreateNewsViewModel struct {
	Title        string `forms:"title" validate:"required"`
	Description  string `forms:"description" validate:"required"`
	Image        string
	CreateUserId string
}
