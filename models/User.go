package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User model structure
type User struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	//NickName string             `json:"nickname,omitempty" bson:"nickname,omitempty" validate:"required"`
	UserName string `json:"username,omitempty" bson:"username,omitempty" validate:"required"`
	Salt     string `bson:"salt,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty" validate:"required"`
}
