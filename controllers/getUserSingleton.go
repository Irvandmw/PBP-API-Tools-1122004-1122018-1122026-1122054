package controllers

import (
	"week9/models"
	"database/sql"
)

var db *sql.DB

func InitializeDB(database *sql.DB) {
	db = database
}

func GetUserByID(userID int) (*models.User, error) {
	query := "SELECT id, name, email, age, points FROM users WHERE id = ?"
	row := db.QueryRow(query, userID)
	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.Points)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
