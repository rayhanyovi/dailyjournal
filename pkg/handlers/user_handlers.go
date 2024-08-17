package handlers

import (
	"dailyjournal/pkg/db"
	"dailyjournal/pkg/models"
	"encoding/json"
	"log"
	"net/http"
)

// CreateUserHandler handles creating a new user
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err := db.Engine.Table("users").Insert(&user)
if err != nil {
    log.Fatalf("Failed to insert user: %v", err)
}
	
	

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUsersHandler handles retrieving all users
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	if err := db.Engine.Find(&users); err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}