package controllers

import (
	"context"
	"fmt"
	"log"
	"mongodbnative/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func insertMovie(movie model.Netflix)  {
	inserted,err := collection.InsertOne(context.Background(),movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted 1 movie to db",inserted.InsertedID)

	
}

func updatemovie(movieId string)  {
	id , err :=primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id":id}
	update := bson.M{"$set":bson.M{"watched":true}}
	res,erro:=collection.UpdateOne(context.Background(),filter,update)
	if erro != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

}

func deleteOne(movieId string)  {
	id , err :=primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id":id}
	count,erro :=collection.DeleteOne(context.Background(),filter)
	if erro != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
	
}

func deleteall()  {
	
	count,err:=collection.DeleteMany(context.Background(),bson.D{{}},nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
	
}

func findallMovies() []primitive.M {
	cur,err := collection.Find(context.Background(),bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M
	for cur.Next(context.Background()){
		var movie bson.M
		erro := cur.Decode(&movie)
		if erro != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)


	}
	
defer cur.Close(context.Background())
return movies

	
}