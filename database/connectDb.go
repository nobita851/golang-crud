package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mongoDb MongoInstance

const dbName = "HR-management"
const mongoURI = "mongodb://localhost:27017/" + dbName

func InitDatabase() (*MongoInstance, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return nil, err
	}

	mongoDb = MongoInstance{
		Client: client,
		Db:     db,
	}
	return &mongoDb, nil
}
