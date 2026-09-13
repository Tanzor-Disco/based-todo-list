package main

import (
	"github.com/Tanzor-Disco/based-todo-list/backend/server"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("godotenv.Load: %v", err)
	}
	URI := os.Getenv("URI")

	server.Run(URI)
}
