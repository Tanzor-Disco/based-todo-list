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

type RegisterRequest struct {
	Username string
	Password string
}

func (s *server) handleRegister(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var registerRequest RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&registerRequest)
	if err != nil {
		log.Printf("handleRegister: json.Decode: %v", err)
		body := newServerResponseBody[any](apperrors.KindErrInvalidJSON, nil)
		sendJSON(w, http.StatusBadRequest, body)
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("handleRegister: bcrypt.GenerateFromPassword: %v", err)
		body := newServerResponseBody[any](apperrors.KindErrInternal, nil)
		sendJSON(w, http.StatusInternalServerError, body)
		return
	}
	userDB := models.NewUser(registerRequest.Username, string(passwordHash))
	userID, err := s.database.CreateUser(ctx, userDB)
	if err != nil {
		log.Printf("handleRegister: %v", err)
		body := newServerResponseBody[any](apperrors.KindErrInternal, nil)
		sendJSON(w, http.StatusInternalServerError, body)
		return
	}

	sessionString, err := auth.CreateSessionString()
	if err != nil {
		log.Printf("handleRegister: %v", err)
		body := newServerResponseBody[any](apperrors.KindErrInternal, nil)
		sendJSON(w, http.StatusInternalServerError, body)
		return
	}

	userSession := models.NewUserSession(userID, sessionString)
	if err = s.database.CreateUserSession(ctx, userSession); err != nil {
		log.Printf("handleRegister: %v", err)
		body := newServerResponseBody[any](apperrors.KindErrInternal, nil)
		sendJSON(w, http.StatusInternalServerError, body)
		return
	}

	cookie := auth.CreateSessionCookie(sessionString)
	http.SetCookie(w, &cookie)

	body := newServerResponseBody[any](apperrors.KindErrNone, nil)
	sendJSON(w, http.StatusOK, body)
}
