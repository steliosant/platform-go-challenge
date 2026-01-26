package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"platform-go-challenge/models"
	"platform-go-challenge/repositories"
)

// AddUser creates a new user
// @Summary Create a new user
// @Description Add a new user to the system
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 201 {object} models.User "User created"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /users [post]
func AddUser(db DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Validate input
		if user.Name == "" {
			http.Error(w, "Name is required", http.StatusBadRequest)
			return
		}

		userID, err := repositories.CreateUser(r.Context(), db.(*sql.DB), user)
		if err != nil {
			println("Error creating user:", err.Error())
			http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		user.ID = userID
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}

// GetUsers retrieves all users
// @Summary Get all users
// @Description Retrieve a list of all users in the system
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} models.User "List of users"
// @Failure 500 {string} string "Internal server error"
// @Router /users [get]
func GetUsers(db DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		users, err := repositories.ListUsers(r.Context(), db.(*sql.DB))
		if err != nil {
			http.Error(w, "Failed to fetch users: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}
}
