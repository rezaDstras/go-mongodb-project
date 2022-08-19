package interfaces

import "golang-api/model/news"

type NewsRepository interface {
	GetNewsList() ([]news.News, error)
	GetNewsById(id string) (news.News, error)
	InsertNews(news news.News) (string, error)
	UpdateNews(news news.News) error
	DeleteNewsById(id string) error
}
