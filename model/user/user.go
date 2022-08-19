package user

import "time"

type User struct {
	//query for get method params
	//json for post method params
	//bson for mongodb fileds
	Id            string    `bson:"_id,omitempty"`
	FirstName     string    `query:"first_name" json:"first_name" bson:"firstName,omitempty"`
	LastFamily    string    `query:"last_name" json:"last_name" bson:"lastName,omitempty"`
	Email         string    `query:"email" json:"email" bson:"email,omitempty"`
	UserName      string    `query:"userName" bson:"userName,omitempty"`
	Password      string    `query:"password" bson:"password,omitempty"`
	RegisterDate  time.Time `query:"registerDate" bson:"registerDate,omitempty"`
	CreatorUserId string    `query:"creatorUserId" bson:"creatorUserId,omitempty"`
	Roles        []string    `bson:"roles,omitempty"`
}
