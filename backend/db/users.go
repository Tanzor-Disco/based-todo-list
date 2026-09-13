package db

import (
	"context"
	"fmt"
	"github.com/Tanzor-Disco/based-todo-list/backend/models"
)

func (db Database) CreateUser(ctx context.Context, user models.User) (models.UserID, error) {
	var userID models.UserID
	err := db.pool.QueryRow(ctx,
		`
	INSERT INTO users(username, password_hash)
	VALUES($1, $2)
	RETURNING id
	`,
		user.Username, user.PasswordHash).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("CreateUser: %w", err)
	}
	return userID, nil
}

func (db Database) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	var user models.User
	err := db.pool.QueryRow(ctx,
		`
	SELECT * FROM users WHERE username = $1
	`,
		username).Scan(&user.ID, &user.Username, &user.PasswordHash)
	if err != nil {
		return user, fmt.Errorf("GetUserByUsername: %w", err)
	}
	return user, nil
}
