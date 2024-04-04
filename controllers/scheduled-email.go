package controllers

import (
	"fmt"
	"log"
	"time"
	"week9/models"

	"github.com/go-co-op/gocron"
)

func NotifyMonthlyPointExpirationEmail() {
	localTime, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		localTime = time.UTC
	}

	s := gocron.NewScheduler(localTime)

	s.Every(1).Month(-7).At("00:01").Do(func() {
		db := connectDB()
		defer db.Close()

		query := "SELECT `email` FROM `users`;"

		rows, err := db.Query(query)
		if err != nil {
			log.Println(err)
			return
		}

		var email string
		var emails []string
		for rows.Next() {
			if err := rows.Scan(&email); err != nil {
				log.Println(err)
				return
			} else {
				emails = append(emails, email)
			}
		}

		subject := "PERINGATAN MASA BERLAKU POIN"
		body := "Masa berlaku poin akan habis pada " + EndOfMonth() + ". Harap menggunakan poin Anda sebelum poin hangus"

		var emailConfig models.EmailConfig
		if err := GetToken(Redis(), "email-config", &emailConfig); err != nil {
			fmt.Print(err)
		}

		for i := 0; i < len(emails); i++ {
			go SendEmail(emailConfig, emails[i], subject, body)
		}
	})
	s.StartBlocking()
}

func SendWeeklyEmail(config models.EmailConfig, recipientEmail string, subject string, body string) error {

	localTime, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		localTime = time.UTC
	}

	s := gocron.NewScheduler(localTime)

	s.Every(1).Monday().At("00:00").Do(func() {
		if err := SendEmail(config, recipientEmail, subject, body); err != nil {
			fmt.Println("Email gagal terkirim")
			return
		}
		fmt.Println("New email has been sent this week")
	})
	s.StartBlocking()
	return nil
}

func SendDailyEmail(config models.EmailConfig, recipientEmail string, subject string, body string) error {

	localTime, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		localTime = time.UTC
	}

	s := gocron.NewScheduler(localTime)

	s.Every(1).Day().At("00:00").Do(func() {
		if err := SendEmail(config, recipientEmail, subject, body); err != nil {
			fmt.Println("Email gagal terkirim")
			return
		}
		fmt.Println("New Email has been sent for today")
	})
	s.StartBlocking()
	return nil
}

func SendEmailByMinute(config models.EmailConfig, recipientEmail string, subject string, body string) error {

	localTime, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		localTime = time.UTC
	}

	s := gocron.NewScheduler(localTime)

	s.Every(1).Minute().Do(func() {
		SendEmail(config, recipientEmail, subject, body)
		fmt.Println("New Email has been sent for this minute")
	})
	s.StartBlocking()
	return nil
}
