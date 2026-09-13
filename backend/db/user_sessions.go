package db

import (
	"context"
	"fmt"

	"github.com/Tanzor-Disco/based-todo-list/backend/models"
)

func (db Database) CreateUserSession(ctx context.Context, userSession models.UserSession) error {
	_, err := db.pool.Exec(ctx,
		`
	INSERT INTO user_sessions(user_id,session_string)
	VALUES($1,$2)
	`,
		userSession.UserID, userSession.SessionString,
	)
	if err != nil {
		return fmt.Errorf("CreateUserSession: %w", err)
	}
	return nil
}

func (db Database) GetUserSessionBySessionString(ctx context.Context, sessionString string) (models.UserSession, error) {
	var user models.UserSession
	err := db.pool.QueryRow(ctx,
		`
	SELECT * FROM user_sessions
	WHERE session_string = $1
	`,
		sessionString).Scan(&user.ID, &user.UserID, &user.SessionString)
	if err != nil {
		return models.UserSession{}, fmt.Errorf("getUserSessionBySessionString: %w", err)
	}
	return user, nil
}
