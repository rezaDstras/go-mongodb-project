package news

import "time"

type News struct {
	//query for get method params
	//json for post method params
	//bson for mongodb fileds

	Id          string `bson:"_id,omitempty"`
	Title       string `query:"title" json:"title" bson:"title,omitempty"`
	Description string `query:"description" json:"description" bson:"description,omitempty"`
	Image       string `query:"image" json:"image" bson:"image,omitempty"`
	CreateDate    time.Time `query:"createDate" bson:"createDate,omitempty"`
	CreatorUserId string `query:"creatorUserId" bson:"creatorUserId,omitempty"`
}
