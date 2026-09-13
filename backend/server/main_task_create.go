package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Tanzor-Disco/based-todo-list/backend/auth"
	"github.com/Tanzor-Disco/based-todo-list/backend/internal/apperrors"
	"github.com/Tanzor-Disco/based-todo-list/backend/models"
)

type TaskRegisterRequest struct {
	Description string
}

func (s server) handleMainTaskCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userSession, err := auth.GetUserSession(r, s.database)
	if err != nil {
		log.Printf("handleMainTaskCreate: %v", err)
		body := newServerResponseBody[any](apperrors.KindErrInvalidSessionString, nil)
		sendJSON(w, http.StatusUnauthorized, body)
		return
	}

	var taskRegisterRequest TaskRegisterRequest
	if err = json.NewDecoder(r.Body).Decode(&taskRegisterRequest); err != nil {
		log.Printf("handleMainTaskCreate: %v", err)
		body := newServerResponseBody[any](apperrors.KindErrInvalidJSON, nil)
		sendJSON(w, http.StatusBadRequest, body)
		return
	}

	taskDB := models.NewTask(userSession.UserID, taskRegisterRequest.Description)
	taskID, err := s.database.CreateTask(ctx, taskDB)
	if err != nil {
		log.Printf("handleMainTaskCreate: %v", err)
		body := newServerResponseBody[any](apperrors.KindErrInvalidJSON, nil)
		sendJSON(w, http.StatusBadRequest, body)
		return
	}

	body := newServerResponseBody(apperrors.KindErrInvalidJSON, []models.TaskID{taskID})
	sendJSON(w, http.StatusOK, body)
}
