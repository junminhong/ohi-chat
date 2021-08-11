package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDb() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load env file")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s", os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGO_HOST"))
	log.Println(mongoURI)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Println(err.Error())
	}
	return client
}

func FindDocByID(mongoDB *mongo.Client, collection *mongo.Collection, objID primitive.ObjectID) primitive.M {
	var (
		err     error
		cursor  *mongo.Cursor
		results []bson.M
		result  primitive.M
	)

	filter := (bson.M{"_id": objID})
	if cursor, err = collection.Find(context.TODO(), filter); err != nil {
		log.Println(err)
	}

	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Println(err)
		}
	}()

	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Println(err)
	}
	for _, resultTmp := range results {
		result = resultTmp
		break
	}
	return result
}

func InsertOneDOC(mongoDB *mongo.Client, collection *mongo.Collection, doc bson.D) interface{} {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, doc)
	if err != nil {
		log.Println(err.Error())
	}
	return res.InsertedID
}
