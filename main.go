package main

import (
	"fmt"
	"log"
	"net/http"

	"week9/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	// CRON JOB
	go controllers.NotifyMonthlyPointExpirationEmail()
	go controllers.PointReset()

	// REDIS
	controllers.SaveToken(controllers.Redis(), "email-config", controllers.NewEmailConfig("smtp.gmail.com", 587, "irvand9999@gmail.com", "ggha yggy gogy lmti"))

	router := mux.NewRouter()
	router.HandleFunc("/user/point/modify", controllers.ModifyPoint).Methods("POST")

	http.Handle("/", router)
	fmt.Println("Membership application listening at 22345")
	log.Println("Membership application listening at 22345")
	log.Fatal(http.ListenAndServe(":22345", router))
}
