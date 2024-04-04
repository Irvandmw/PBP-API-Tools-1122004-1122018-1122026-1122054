package main

import (
	"fmt"
	"log"
	"net/http"
	"week9/controllers"

	"github.com/gorilla/mux"
)

func main() {

	// testing redis
	controllers.RedisClient()

	router := mux.NewRouter()
	router.HandleFunc("/user/point/modify", controllers.ModifyPoint).Methods("POST")

	http.Handle("/", router)
	fmt.Println("Membership application listening at 22345")
	log.Println("Membership application listening at 22345")
	log.Fatal(http.ListenAndServe(":22345", router))
}
