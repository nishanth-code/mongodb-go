package main

import (
	"fmt"
	"log"
	"mongodbnative/router"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("server starting...")
	log.Fatal(http.ListenAndServe(":5000",r))
	fmt.Println("sever started")


}
