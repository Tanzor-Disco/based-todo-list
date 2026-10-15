package main

import (
	"github.com/Tanzor-Disco/based-todo-list/backend/db/migrations"
	"github.com/Tanzor-Disco/based-todo-list/backend/server"
	"github.com/joho/godotenv"
	"log"
	"os"
)

//testing

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("godotenv.Load: %v", err)
	}
	URI := os.Getenv("URI")

	err = migrations.Update(URI)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatalf("server error: %v", server.Run(URI))
}
