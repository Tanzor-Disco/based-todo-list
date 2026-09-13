package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Tanzor-Disco/based-todo-list/backend/db"
)

type server struct {
	database db.Database
}

func newServer(database db.Database) server {
	return server{
		database: database,
	}
}

type serverResponseBody[T any] struct {
	ErrorKind string `json:"error_kind"`
	Data      []T    `json:"data"`
}

func newServerResponseBody[T any](errorKind string, data []T) serverResponseBody[T] {
	return serverResponseBody[T]{
		ErrorKind: errorKind,
		Data:      data,
	}
}

func sendJSON[T any](w http.ResponseWriter, code int, body serverResponseBody[T]) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Println(err)
	}
	_, err = w.Write(bodyBytes)
	if err != nil {
		log.Println(err)
	}
}

func Run(URI string) error {
	database, err := db.Connect(URI)
	if err != nil {
		return fmt.Errorf("Run: %w", err)
	}
	server := newServer(database)

	http.HandleFunc("/api/register", server.handleRegister)
	http.HandleFunc("/api/main/task/create", server.handleMainTaskCreate)
	http.HandleFunc("/api/main/tasks", server.handleMainTasks)
	http.HandleFunc("/api/login", server.handleLogin)
	http.HandleFunc("/", handleStatic)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		return fmt.Errorf("Run: ListenAndServe: %w", err)
	}
	return nil
}
