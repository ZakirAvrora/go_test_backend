package store

import (
	"ZakirAvrora/go_test_backend/service1/internal/model"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ErrUserExists = errors.New("the user already exists")

type MongoStore struct {
	collection *mongo.Collection
	ctx        context.Context
}

func New(collection *mongo.Collection, ctx context.Context) *MongoStore {
	return &MongoStore{collection: collection,
		ctx: ctx}
}

func (s *MongoStore) CreateUser(user model.DbUser) error {
	_, err := s.collection.InsertOne(s.ctx, user)

	if err != nil {
		if err1, ok := err.(mongo.WriteException); ok && err1.WriteErrors[0].Code == 11000 {
			return fmt.Errorf("user with the same email already exists: %w", ErrUserExists)
		}
	}

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}

	if _, err := s.collection.Indexes().CreateOne(s.ctx, index); err != nil {
		return fmt.Errorf("could not create index for email: %w", err)
	}

	return nil
}

func (s *MongoStore) GetUser(email string) (*model.DbUser, error) {
	var user *model.DbUser

	if err := s.collection.FindOne(s.ctx, bson.M{"email": email}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no user with that email is found: %w", err)
		}

		return nil, fmt.Errorf("error in getting user by email: %w", err)
	}

	return user, nil
}
