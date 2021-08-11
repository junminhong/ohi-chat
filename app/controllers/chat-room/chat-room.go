package chatroom

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/junminhong/ohi-chat/config/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(c *gin.Context) {

}

func List(c *gin.Context) {
	mongoDB := database.InitMongoDb()
	collection := mongoDB.Database("chat_room").Collection("test")
	objID, err := primitive.ObjectIDFromHex("6113ce4220c235223eb3b6fe")
	if err != nil {
		log.Println(err.Error())
	}
	resultDOC := database.FindDocByID(mongoDB, collection, objID)
	log.Println(resultDOC)
	doc := bson.D{{"name", "hello"}, {"value", "ass"}}
	docID := database.InsertOneDOC(mongoDB, collection, doc)
	log.Println(docID)
}

func Edit(c *gin.Context) {

}

func Delete(c *gin.Context) {

}

func Join(c *gin.Context) {

}

func Leave(c *gin.Context) {

}
