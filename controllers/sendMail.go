package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"week9/models"

	gomail "gopkg.in/mail.v2"
)

var db *sql.DB

func InitializeDB(database *sql.DB) {
	db = database
}

func SendEmail(config *models.EmailConfig, recipientEmail string, subject string, body string) error {
	email := gomail.NewMessage()
	email.SetHeader("From", config.SenderEmail)
	email.SetHeader("To", recipientEmail)
	email.SetHeader("Subject", subject)
	email.SetBody("text/html", body)
	dialer := gomail.NewDialer(config.Host, config.Port, config.SenderEmail, config.SenderPassword)

	if err := dialer.DialAndSend(email); err != nil {
		return err
	}

	return nil
}

func PenambahanPoin(config *models.EmailConfig, w http.ResponseWriter, r *http.Request) error {
	db := connect()
	defer db.Close()
	var userID = r.URL.Query().Get("id")
	poinStr := r.URL.Query().Get("poin")
	poin, err := strconv.Atoi(poinStr)

	if err != nil {
		return fmt.Errorf("%s bukan poin yang valid", poinStr)
	}

	query := "SELECT id, name, email, age, points FROM users WHERE id = ?"
	row := db.QueryRow(query, userID)
	var user models.User
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.Points)
	if err != nil {
		return err
	}

	tempPoin := user.Points + poin

	query = "UPDATE users SET points = ? WHERE id = ?"
	_, err = db.Exec(query, tempPoin, userID)
	if err != nil {
		return err
	}

	subject := "Penambahan Poin"

	body := fmt.Sprintf("Halo %s,\n\nAnda telah menerima penambahan %d poin. Total poin saat ini: %d\n\nTerima kasih.", user.Name, poin, tempPoin)

	// Pembuatan Email
	email := gomail.NewMessage()
	email.SetHeader("From", config.SenderEmail)
	email.SetHeader("To", user.Email)
	email.SetHeader("Subject", subject)
	email.SetBody("text/html", body)
	dialer := gomail.NewDialer(config.Host, config.Port, config.SenderEmail, config.SenderPassword)

	fmt.Print(config.SenderEmail)
	if err := dialer.DialAndSend(email); err != nil {
		fmt.Print("AnjayMabar2")
		return err
	}
	return nil
}

func PenguranganPoin(config *models.EmailConfig, w http.ResponseWriter, r *http.Request) error {
	db := connect()
	defer db.Close()
	var userID = r.URL.Query().Get("id")
	poinStr := r.URL.Query().Get("poin")
	poin, err := strconv.Atoi(poinStr)
	if err != nil {
		return fmt.Errorf("%s bukan poin yang valid", poinStr)
	}

	query := "SELECT id, name, email, age, points FROM users WHERE id = ?"
	row := db.QueryRow(query, userID)
	var user models.User
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.Points)
	if err != nil {
		return err
	}

	subject := "Pengurangan Poin"

	var body string
	if user.Points-poin < 0 {
		body = fmt.Sprintf("Halo %s,\n\nAnda tidak memiliki poin yang cukup untuk melakukan pengurangan sebanyak %d poin. Total poin saat ini: %d\n\nTerima kasih.", user.Name, poin, user.Points)
	} else {
		tempPoin := user.Points - poin
		body = fmt.Sprintf("Halo %s,\n\nAnda telah menerima pennguragan %d poin. Total poin saat ini: %d\n\nTerima kasih.", user.Name, poin, tempPoin)
		//update total poin
		query = "UPDATE users SET points = ? WHERE id = ?"
		_, err = db.Exec(query, tempPoin, userID)
		if err != nil {
			return err
		}
	}

	// Pembuatan Email
	email := gomail.NewMessage()
	email.SetHeader("From", config.SenderEmail)
	email.SetHeader("To", user.Email)
	email.SetHeader("Subject", subject)
	email.SetBody("text/html", body)
	dialer := gomail.NewDialer(config.Host, config.Port, config.SenderEmail, config.SenderPassword)

	if err := dialer.DialAndSend(email); err != nil {
		return err
	}
	return nil

}
