package interfaces

import (
	"golang-api/model/news"
	newsViewModel "golang-api/viewModel/news"
)

type NewsService interface {
	GetNewsList() ([]news.News, error)
	CreateNews(newsInput newsViewModel.CreateNewsViewModel) (string, error)

}
