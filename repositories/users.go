package repositories

import (
	"context"
	"database/sql"

	"platform-go-challenge/models"
)

func CreateUser(
	ctx context.Context,
	db *sql.DB,
	user models.User,
) (string, error) {
	query := `
	INSERT INTO users (id, name)
	VALUES ($1, $2)
	RETURNING id;
	`

	var userID string
	err := db.QueryRowContext(ctx, query, user.ID, user.Name).Scan(&userID)
	if err != nil {
		return "", err
	}

	return userID, nil
}

func GetUserByID(
	ctx context.Context,
	db *sql.DB,
	userID string,
) (*models.User, error) {
	query := `
	SELECT id, name
	FROM users
	WHERE id = $1;
	`

	var user models.User
	err := db.QueryRowContext(ctx, query, userID).Scan(&user.ID, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
