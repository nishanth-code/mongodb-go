package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongodbnative/model"
	"net/http"

	"github.com/gorilla/mux"
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

func GetallMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	allMovies := findallMovies()
	json.NewEncoder(w).Encode(allMovies)

}

func CreateMovie(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var movie model.Netflix
	_=json.NewDecoder(r.Body).Decode(&movie)
	insertMovie(movie)
	json.NewEncoder(w).Encode(movie)

	
}

func MarkMovieaswatched(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")

	params := mux.Vars(r)
	updatemovie(params["id"])
	json.NewEncoder(w).Encode("updated sucessfully")


	
}

func DeleteOne(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	params := mux.Vars(r)
	deleteOne(params["id"])
	json.NewEncoder(w).Encode("deleted 1 sucessfully")


	
}

func DeleteAll(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	deleteall()
	json.NewEncoder(w).Encode("deleted all sucessfully")


}

