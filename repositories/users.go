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

func ListUsers(
	ctx context.Context,
	db *sql.DB,
) ([]models.User, error) {
	query := `
	SELECT id, name
	FROM users
	ORDER BY id;
	`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
