package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-api/database"
	"golang-api/model/news"
	newsRepoInterface "golang-api/repository/interfaces"
	"log"
)

type newsRepository struct {
	Db database.Db
}

func NewNewsRepository() newsRepoInterface.NewsRepository {
	db, err := database.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	return newsRepository{
		Db: db,
	}
}

func (u newsRepository) GetNewsList() ([]news.News, error) {
	newsCollection := u.Db.GetNewsCollection()
	cursor, err := newsCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var Newss []news.News
	err = cursor.All(context.TODO(), &Newss)
	if err != nil {
		return nil, err
	}
	return Newss, nil
}

func (u newsRepository) GetNewsById(id string) (news.News, error) {
	//convert id string to objectid in mongodb
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return news.News{}, err
	}
	NewsCollection := u.Db.GetNewsCollection()
	var NewsObject news.News

	err = NewsCollection.FindOne(context.TODO(), bson.D{
		{"_id", objectId},
	}).Decode(&NewsObject)
	if err != nil {
		return news.News{}, err
	}
	return NewsObject, nil
}

func (u newsRepository) InsertNews(News news.News) (string, error) {
	NewsCollection := u.Db.GetNewsCollection()
	res, err := NewsCollection.InsertOne(context.TODO(), News)
	if err != nil {
		return "", err
	}
	// res pass InsertedId
	//convert inserted id to object id
	objectId := res.InsertedID.(primitive.ObjectID).Hex()

	return objectId, nil

}

func (u newsRepository) UpdateNews(News news.News) error {
	//convert id string to objectid in mongodb
	objectId, err := primitive.ObjectIDFromHex(News.Id)
	if err != nil {
		return err
	}
	//for avoiding update _id
	News.Id = ""
	NewsCollection := u.Db.GetNewsCollection()
	_, err = NewsCollection.UpdateOne(context.TODO(), bson.D{{"_id", objectId}}, bson.D{{"$set", News}})
	if err != nil {
		return err
	}
	return nil

}

func (u newsRepository) DeleteNewsById(id string) error {
	//convert id string to objectid in mongodb
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	NewsCollection := u.Db.GetNewsCollection()
	_, err = NewsCollection.DeleteOne(context.TODO(), bson.D{{"_id", objectId}})
	if err != nil {
		return err
	}
	return nil
}
