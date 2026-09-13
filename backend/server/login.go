package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Tanzor-Disco/based-todo-list/backend/auth"
	"github.com/Tanzor-Disco/based-todo-list/backend/internal/apperrors"
	"github.com/Tanzor-Disco/based-todo-list/backend/models"
	"golang.org/x/crypto/bcrypt"
)

type loginRequest struct {
	Username string
	Password string
}

func (s server) handleLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var loginRequest loginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		log.Printf("handleLogin: %v", err)
		body := newServerResponseBody[any](apperrors.KindErrInvalidJSON, nil)
		sendJSON(w, http.StatusBadRequest, body)
		return
	}

	user, err := s.database.GetUserByUsername(ctx, loginRequest.Username)
	if err != nil {
		log.Printf("handleLogin: %v", err)
		body := newServerResponseBody[any](apperrors.KindErrInvalidCredentials, nil)
		sendJSON(w, http.StatusUnauthorized, body)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginRequest.Password))
	if err != nil {
		log.Printf("handleLogin: %v", err)
		body := newServerResponseBody[any](apperrors.KindErrInvalidCredentials, nil)
		sendJSON(w, http.StatusUnauthorized, body)
		return
	}
	sessionString, err := auth.CreateSessionString()
	if err != nil {
		log.Printf("handleLogin: %v", err)
		body := newServerResponseBody[any](apperrors.KindErrInternal, nil)
		sendJSON(w, http.StatusInternalServerError, body)
		return
	}
	userSession := models.NewUserSession(user.ID, sessionString)
	if err = s.database.CreateUserSession(ctx, userSession); err != nil {
		log.Printf("handleLogin: %v", err)
		body := newServerResponseBody[any](apperrors.KindErrInternal, nil)
		sendJSON(w, http.StatusInternalServerError, body)
		return
	}
	cookie := auth.CreateSessionCookie(sessionString)
	http.SetCookie(w, &cookie)
	body := newServerResponseBody[any](apperrors.KindErrNone, nil)
	sendJSON(w, http.StatusOK, body)
}
