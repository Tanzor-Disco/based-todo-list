package auth

import (
	"fmt"
	"net/http"

	"github.com/Tanzor-Disco/based-todo-list/backend/db"
	"github.com/Tanzor-Disco/based-todo-list/backend/models"
)

func GetUserSession(r *http.Request, db db.Database) (models.UserSession, error) {
	cookie, err := r.Cookie("session_string")
	if err != nil {
		return models.UserSession{}, fmt.Errorf("getUserSession: r.Cookie: %w", err)
	}
	sessionString := cookie.Value
	userSession, err := db.GetUserSessionBySessionString(r.Context(), sessionString)
	if err != nil {
		return models.UserSession{}, fmt.Errorf("getUserSession: %w", err)
	}
	return userSession, nil
}
