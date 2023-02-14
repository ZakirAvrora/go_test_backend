package db

import (
	"ZakirAvrora/go_test_backend/service1/config"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Init(conf *config.DatabaseConfig) (*mongo.Client, error) {

	// confStr := fmt.Sprintf("mongodb://%s:%s@%s:%s/?ssl=false&maxPoolSize=20&w=majority",
	// conf.User, conf.Password, conf.Host, conf.Port)

	confStr := "mongodb://mongo:27017"

	client, err := mongo.Connect(context.TODO(),
		options.Client().ApplyURI(confStr))

	if err != nil {
		return nil, fmt.Errorf("connection to db failed: %w", err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, fmt.Errorf("cannot ping to db  %v: %w", confStr, err)
	}

	return client, nil
}
