package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"mallekoppie/script-workflow-engine/server/models"
)

const (
	collectionNameTestGroups      string = "scripts"
	collectionNameTestCollections string = "workflows"
	fieldNameId                   string = "_id"
	FieldName                     string = "groupid"
)

var (
	mongoClient   *mongo.Client
	mongoDBName   string
	serviceConfig models.ServiceConfig

	ErrUpdateCountWrong                 = errors.New("Incorrect update count")
	ErrNoGroupExistsForTestCollectionId = errors.New("No group exists for ID")
)

func init() {
	connectToMongo()

	serviceConfig, _ = GetConfig()
	mongoDBName = serviceConfig.MongoDBName
}

func connectToMongo() {
	config, err := GetConfig()
	if err != nil {
		log.Panicln("Unable to read service config to get mongoDB details: ", err.Error())
		os.Exit(1)
	}

	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%v:%v", config.MongoDBHost, config.MongoDBPort))
	mongoClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Panicln("Unable to connect to mongo:", err.Error())
		os.Exit(1)
	}

	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Panicln("Unable to ping mongodb: ", err.Error())
		os.Exit(1)
	}

	log.Println("Mongo connection successfull!")
}
