package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"platform-go-challenge/models"
	"platform-go-challenge/repositories"
)

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
