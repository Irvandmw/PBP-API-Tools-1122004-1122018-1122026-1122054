package main

import (
	"fmt"
	"log"
	"net/http"

	"week9/controllers"
	"week9/models"

	_ "github.com/go-sql-driver/mysql"
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

	// config email sender sementara untuk bisa send email
	config := models.NewEmailConfig(
		"smtp.gmail.com",
		587,
		"irvand9999@gmail.com",
		"ggha yggy gogy lmti",
	)
	
	//END POINTS
	router.HandleFunc("/user/tambahPoin", func(w http.ResponseWriter, r *http.Request) {
		err := controllers.PenambahanPoin(config, w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}).Methods("POST")
	router.HandleFunc("/user/kurangPoin", func(w http.ResponseWriter, r *http.Request) {
		err := controllers.PenguranganPoin(config, w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}).Methods("POST")
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}

// branchWilliam
	//Ini bagian yang bisa dimodifikasi (line 20-22)
	recipientEmail := "if-22004@students.ithb.ac.id"
	subject := "Test Go Mail doang"
	body := "Hello <b>Irvan</b> ini adalah sebuah test email dari Gomail"

	if err := controllers.SendEmail(config, recipientEmail, subject, body); err != nil {
		fmt.Println("Gagal mengirim email:", err)
		return
	}

	fmt.Println("Email berhasil dikirim!")

	// testing redis
	// controllers.RedisClient()

	// SCHEDULING EMAIL GOCRON
	controllers.SendEmailByMinute(config, recipientEmail, "PESAN MENIT INI", "Ini pesan yang muncul setiap menit")
	controllers.SendDailyEmail(config, recipientEmail, "PESAN HARI INI", "Ini pesan yang muncul setiap hari")
	controllers.SendWeeklyEmail(config, recipientEmail, "PESAN MINGGU INI", "Ini pesan yang muncul setiap minggu")
	controllers.SendMonthlyEmail(config, recipientEmail, "PESAN BULAN INI", "Ini pesan yang muncul setiap bulan")
}
