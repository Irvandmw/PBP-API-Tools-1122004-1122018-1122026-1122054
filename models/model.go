package models

type EmailConfig struct {
    Host           string
    Port           int
    SenderEmail    string
    SenderPassword string
}

func NewEmailConfig(host string, port int, senderEmail, senderPassword string) *EmailConfig {
    return &EmailConfig{
        Host:           host,
        Port:           port,
        SenderEmail:    senderEmail,
        SenderPassword: senderPassword,
    }
}
