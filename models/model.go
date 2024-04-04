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
