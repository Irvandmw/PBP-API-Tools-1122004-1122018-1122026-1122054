package models

type EmailConfig struct {
	Host           string
	Port           int
	SenderEmail    string
	SenderPassword string
}

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Points int    `json:"points"`
}

// Request struct section
type ModifyPointRequest struct {
	UserID int `json:"id_user"`
	Amount int `json:"point"`
}

// Response struct section
type BasicResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Points int    `json:"points"`
}

func NewEmailConfig(host string, port int, senderEmail, senderPassword string) *EmailConfig {
	return &EmailConfig{
		Host:           host,
		Port:           port,
		SenderEmail:    senderEmail,
		SenderPassword: senderPassword,
	}
type User struct {
	ID     int
	Name   string
	Email  string
	Age    int
	Points int
}

// Request struct section
type ModifyPointRequest struct {
	UserID int `json:"id_user"`
	Amount int `json:"point"`
}

// Response struct section
type BasicResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewUser(id int, name string, email string, age int, points int) *User {
	return &User{
		ID:     id,
		Name:   name,
		Email:  email,
		Age:    age,
		Points: points,
	}
}
