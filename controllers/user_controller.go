package controllers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"week9/models"
)

func GetUserByID(userID int) (*models.User, error) {
	db := connectDB()
	defer db.Close()

	query := "SELECT id, name, email, age, points FROM users WHERE id = ?"
	row := db.QueryRow(query, userID)
	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.Points)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func ModifyPoint(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var request models.ModifyPointRequest
	err := decoder.Decode(&request)
	if err != nil {
		sendErrorResponse(w, 400, "empty request body")
		return
	}
	id_user := request.UserID
	if id_user <= 0 {
		sendErrorResponse(w, 400, "invalid id_user")
		return
	}
	amount := request.Amount
	if amount <= 0 {
		sendErrorResponse(w, 400, "invalid point")
		return
	}

	// Update data ke db
	user, err := GetUserByID(request.UserID)
	if err != nil {
		sendErrorResponse(w, 400, "user not found")
	}

	user.Points += request.Amount
	query := "UPDATE `users` SET `points`=? WHERE `id`=?;"
	_, err_update := db.Exec(query, user.Points, user.ID)
	if err_update != nil {
		sendErrorResponse(w, 500, "internal server error")
	}

	query_insert := "INSERT INTO `point_log`(`timestamp`, `id_user`, `point`) VALUE (now(), ?, ?);"
	_, err_insert_log := db.Exec(query_insert, user.ID, request.Amount)
	if err_insert_log != nil {
		sendErrorResponse(w, 500, "unable to print point log")
	}

	var emailConfig models.EmailConfig
	if err := GetToken(Redis(), "email-config", &emailConfig); err != nil {
		fmt.Print(err)
	}

	fmt.Print(emailConfig.Host)

	if request.Amount > 0 {
		go KirimPenambahanPoin(emailConfig, user, request.Amount)
	} else {
		go KirimPenguranganPoin(emailConfig, user, int(math.Abs(float64(request.Amount))))
	}
	sendSuccessResponse(w, "success")
}

func sendSuccessResponse(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application-json")
	var response models.BasicResponse
	response.Status = 200
	response.Message = message
	json.NewEncoder(w).Encode(response)
}
func sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application-json")
	var response models.ErrorResponse
	response.Status = statusCode
	response.Message = message
	json.NewEncoder(w).Encode(response)
}
