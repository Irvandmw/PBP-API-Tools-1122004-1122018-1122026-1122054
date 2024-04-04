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
}
	// config email sender sementara untuk bisa send email
	config := models.NewEmailConfig(
		"smtp.gmail.com",
		587,
		"irvand9999@gmail.com",
		"ggha yggy gogy lmti",
	)
	router := mux.NewRouter()
	
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

//DONT MIND ANYTHING BELOW THIS LINE

// Loop:
//     for{
//         fmt.Print("Masukkan ID pengguna (0 untuk keluar): ")
//         var userID int
//         if userID==0 {
//             fmt.Println("Sampai Jumpa!")
//             os.Exit(0)
//         }
//         user, err := controllers.GetUserByID(userID)
//         if err != nil {
//             fmt.Println("Failed to get user data:", err)
//             continue Loop
//         }

//         var status = true
//         for status {
//             fmt.Println("Pilih fungsi yang ingin dijalankan:")
//             fmt.Println("1. Penggunaan poin")
//             fmt.Println("2. Penambahan poin")
//             fmt.Println("3. LogOut")
//             fmt.Print("Masukkan nomor fungsi: ")

//             reader := bufio.NewReader(os.Stdin)
//             input, _ := reader.ReadString('\n')
//             input = input[:len(input)-1]

//             num, err := strconv.Atoi(input)
//             if err != nil {
//                 fmt.Println("Input tidak valid:", err)
//                 return
//             }

//             switch num {
//             case 1:
//                 controllers.PenambahanPoin(config, user, 10)
//             case 2:
//                 controllers.PenguranganPoin(config, user, 5)
//             case 3:
//                 status = false
//             default:
//                 fmt.Println("Fungsi tidak valid")
//             }
//         }
//     }
// }

// 	//Ini bagian yang bisa dimodifikasi (line 20-22)
//     recipientEmail := "irvand9999@gmail.com"
//     subject := "Test Go Mail doang"
//     body := "Hello <b>Irvan</b> ini adalah sebuah test email dari Gomail"

//     if err := controllers.SendEmail(config, recipientEmail, subject, body); err != nil {
//         fmt.Println("Gagal mengirim email:", err)
//         return
//     }

//     fmt.Println("Email berhasil dikirim!")
// }
