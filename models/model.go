package models

type EmailConfig struct {
    Host           string
    Port           int
    SenderEmail    string
    SenderPassword string
}

type User struct {
    ID    int
    Name  string
    Email string
    Age   int
    Points int
}

func NewEmailConfig(host string, port int, senderEmail, senderPassword string) *EmailConfig {
    return &EmailConfig{
        Host:           host,
        Port:           port,
        SenderEmail:    senderEmail,
        SenderPassword: senderPassword,
    }
}

func NewUser(id int, name string,email string, age int, points int) *User {
    return &User{
        ID:    id,
        Name:  name,
        Email: email,
        Age:   age,
        Points: points,
    }
}
