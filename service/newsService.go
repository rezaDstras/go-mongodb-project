package service

import (
	"golang-api/model/news"
	"golang-api/repository"
	newsServiceInterface "golang-api/service/interfaces"
	newsViewModel "golang-api/viewModel/news"
	"time"
)

type newsService struct {
}

func NewnewsService() newsServiceInterface.NewsService {
	return newsService{}
}

func (n newsService) GetNewsList() ([]news.News, error) {
	newsRepository := repository.NewNewsRepository()
	newsList , err := newsRepository.GetNewsList()
	return newsList,err
}

func (n newsService)CreateNews(newsInput newsViewModel.CreateNewsViewModel) (string, error) {

	newsEntity := news.News{
		Title:         newsInput.Title,
		Description:   newsInput.Description,
		Image: newsInput.Image,
		CreateDate: time.Now(),
	}


	newsRepository := repository.NewNewsRepository()
	newsInsert , err := newsRepository.InsertNews(newsEntity)

	return newsInsert, err
}


