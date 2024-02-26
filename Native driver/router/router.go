package router

import (
	"mongodbnative/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies",controllers.GetallMovies).Methods("GET")
	router.HandleFunc("/api/movie",controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}",controllers.MarkMovieaswatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}",controllers.DeleteOne).Methods("DELETE")
	router.HandleFunc("/api/movies",controllers.DeleteAll).Methods("DELETE")
	
	


	return router

}