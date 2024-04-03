package main

import (
	"week9/controllers"
)

func main() {
	// config sementara untuk bisa send email
	// config := models.NewEmailConfig(
	// 	"smtp.gmail.com",
	// 	587,
	// 	"irvand9999@gmail.com",
	// 	"ggha yggy gogy lmti",
	// )

	// //Ini bagian yang bisa dimodifikasi (line 20-22)
	// recipientEmail := "irvand9999@gmail.com"
	// subject := "Test Go Mail doang"
	// body := "Hello <b>Irvan</b> ini adalah sebuah test email dari Gomail"

	// if err := controllers.SendEmail(config, recipientEmail, subject, body); err != nil {
	// 	fmt.Println("Gagal mengirim email:", err)
	// 	return
	// }

	// fmt.Println("Email berhasil dikirim!")

	// testing redis
	controllers.RedisClient()
}
