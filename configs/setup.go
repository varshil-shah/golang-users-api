package configs

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDatabase() *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoUrl, er := GetEnvVariable("MONGO_URL")

	if er != nil {
		log.Fatal(fmt.Errorf(er.Error()))
	}

	opts := options.Client().ApplyURI(mongoUrl).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected successfully!")

	return client
}

var DB *mongo.Client = ConnectDatabase()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("golang-api").Collection(collectionName)
	return collection
}
