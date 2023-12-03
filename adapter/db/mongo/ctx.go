package mongo

import (
	"context"
	"fmt"
	"log"

	"myTestWithMultiStage/env"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func Connect() {
	connectionString := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/",
		env.Var.MongoUsername,
		env.Var.MongoPassword,
		env.Var.MongoHost,
		env.Var.MongoPort,
	)

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Couldn't connect to MongoDB:", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Couldn't ping MongoDB:", err)
	}

	DB = client
}
