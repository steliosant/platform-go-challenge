package handlers

import (
	"net/http"
	"strings"
)

func FavouritesRouter(db DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Expected paths:
		// /users/{userID}/favourites
		// /users/{userID}/favourites/{assetID}

		parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

		// Minimum: users/{id}/favourites
		if len(parts) < 3 || parts[0] != "users" || parts[2] != "favourites" {
			http.NotFound(w, r)
			return
		}

		switch r.Method {

		case http.MethodGet:
			// GET /users/{id}/favourites
			if len(parts) == 3 {
				println("Get lists of favorite assets for a user")
				GetUserFavourites(db)(w, r)
				return
			}

		case http.MethodPost:
			// POST /users/{id}/favourites
			if len(parts) == 3 {
				println("Add new favorite asset to a user user")
				AddFavourite(db)(w, r)
				return
			}

		case http.MethodPatch:
			// PATCH /users/{id}/favourites/{assetID}
			if len(parts) == 4 {
				println("Update favorite assets for a user user")
				UpdateFavourite(db)(w, r)
				return
			}

		case http.MethodDelete:
			// DELETE /users/{id}/favourites/{assetID}
			if len(parts) == 4 {
				println("Remove asset for a user user")
				RemoveFavourite(db)(w, r)
				return
			}
		}

		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
