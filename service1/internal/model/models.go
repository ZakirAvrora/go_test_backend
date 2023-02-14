package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type PostReqUser struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type Salt struct {
	Salt string `json:"salt" binding:"required"`
}

type DbUser struct {
	Id       primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Salt     string             `json:"salt,omitempty" bson:"salt,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
}
