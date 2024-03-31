package controllers

import (
    "week9/models"
    gomail "gopkg.in/mail.v2"
)

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


