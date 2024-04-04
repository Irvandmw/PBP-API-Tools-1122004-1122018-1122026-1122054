package controllers

import (
	"fmt"
	"time"
	"week9/models"

	"github.com/go-co-op/gocron"
)

func SendMonthlyEmail(config models.EmailConfig, recipientEmail string, subject string, body string) error {

	localTime, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		localTime = time.UTC
	}

	s := gocron.NewScheduler(localTime)

	s.Every(1).Month().At("00:00").Do(func() {
		if err := SendEmail(config, recipientEmail, subject, body); err != nil {
			fmt.Println("Email gagal terkirim")
			return
		}
		fmt.Println("New Email has been sent this month")
	})
	s.StartBlocking()
	return nil
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
