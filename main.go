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

	controllers.SaveToken(controllers.Redis(), "email-config", controllers.NewEmailConfig("smtp.gmail.com", 587, "irvand9999@gmail.com", "ggha yggy gogy lmti"))

	router := mux.NewRouter()
	router.HandleFunc("/user/point/modify", controllers.ModifyPoint).Methods("POST")

	http.Handle("/", router)
	fmt.Println("Membership application listening at 22345")
	log.Println("Membership application listening at 22345")
	log.Fatal(http.ListenAndServe(":22345", router))

	// SCHEDULING EMAIL GOCRON
	// controllers.SendEmailByMinute(config, recipientEmail, "PESAN MENIT INI", "Ini pesan yang muncul setiap menit")
	// controllers.SendDailyEmail(config, recipientEmail, "PESAN HARI INI", "Ini pesan yang muncul setiap hari")
	// controllers.SendWeeklyEmail(config, recipientEmail, "PESAN MINGGU INI", "Ini pesan yang muncul setiap minggu")
	// controllers.SendMonthlyEmail(config, recipientEmail, "PESAN BULAN INI", "Ini pesan yang muncul setiap bulan")
}
