package controllers

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ConnectionString = "mongodb+srv://nishanth:nish1234@cluster0.bbodeek.mongodb.net/"
const colname = "watchlist"

var collection *mongo.Collection


func init()  {
	Clientoptions := options.Client().ApplyURI(ConnectionString)

	Client,err := mongo.Connect(context.TODO(),Clientoptions)

	if err != nil {
		panic(err)
	}
	collection = Client.Database("netflix").Collection(colname)
	fmt.Println("collect instance ready")
	
}