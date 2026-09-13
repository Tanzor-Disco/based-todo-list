package server

import (
	"log"
	"net/http"

	"github.com/Tanzor-Disco/based-todo-list/backend/auth"
	"github.com/Tanzor-Disco/based-todo-list/backend/internal/apperrors"
)

func (s server) handleMainTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userSession, err := auth.GetUserSession(r, s.database)
	if err != nil {
		log.Printf("handleMain: %v", err)
		body := newServerResponseBody[any](apperrors.KindErrInvalidSessionString, nil)
		sendJSON(w, http.StatusUnauthorized, body)
		return
	}

	tasks, err := s.database.GetTasksByUserID(ctx, userSession.UserID)
	if err != nil {
		log.Printf("handleMain: %v", err)
		body := newServerResponseBody[any](apperrors.KindErrInternal, nil)
		sendJSON(w, http.StatusInternalServerError, body)
		return
	}

	body := newServerResponseBody(apperrors.KindErrNone, tasks)
	sendJSON(w, http.StatusOK, body)
}
