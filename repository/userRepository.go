package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-api/database"
	"golang-api/model/user"
	userRepoInterface "golang-api/repository/interfaces"
	"log"
)

type userRepository struct {
	Db database.Db
}

func NewUserRepository() userRepoInterface.UserRepository{
	db, err := database.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	return userRepository{
		Db: db,
	}
}

func (u userRepository) GetUserList() ([]user.User, error) {
	userCollection := u.Db.GetUserCollection()
	cursor, err := userCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var users []user.User
	err = cursor.All(context.TODO(), &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u userRepository) GetUserById(id string) (user.User, error) {
	//convert id string to objectid in mongodb
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user.User{}, err
	}
	userCollection := u.Db.GetUserCollection()
	var userObject user.User

	err = userCollection.FindOne(context.TODO(), bson.D{
		{"_id", objectId},
	}).Decode(&userObject)
	if err != nil {
		return user.User{}, err
	}
	return userObject, nil
}

func (u userRepository) GetUserByusernameAndPassword(username , password string) (user.User, error) {
	//convert id string to objectid in mongodb
	userCollection := u.Db.GetUserCollection()
	var userObject user.User

	err := userCollection.FindOne(context.TODO(), bson.D{
		{"userName", username},
		{"password", password},
	}).Decode(&userObject)
	if err != nil {
		return user.User{}, err
	}
	return userObject, nil
}

func (u userRepository) InsertUser(user user.User) (string, error) {
	userCollection := u.Db.GetUserCollection()
	res, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return "", err
	}
	// res pass InsertedId
	//convert inserted id to object id
	objectId := res.InsertedID.(primitive.ObjectID).Hex()

	return objectId, nil

}

func (u userRepository) UpdateUser(user user.User) error {
	//convert id string to objectid in mongodb
	objectId, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return err
	}
	//for avoiding update _id
	user.Id = ""
	userCollection := u.Db.GetUserCollection()
	_, err = userCollection.UpdateOne(context.TODO(), bson.D{{"_id", objectId}}, bson.D{{"$set", user}})
	if err != nil {
		return err
	}
	return nil

}

func (u userRepository)DeleteUserById(id string)error  {
	//convert id string to objectid in mongodb
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return  err
	}
	userCollection := u.Db.GetUserCollection()
	_,err =userCollection.DeleteOne(context.TODO(),bson.D{{"_id",objectId}})
	if err != nil {
		return  err
	}
	return nil
}
