package main

import (
	"log"

	"github.com/Kilril312/users-service/internal/database"
	"github.com/Kilril312/users-service/internal/user"
)

func main() {
	db := database.InitDB()

	repo := user.NewRepository(db)
	svc := user.NewService(repo)

	if err := transport.RunGRPC(svc); err != nil {
		log.Fatal("gRPC сервер завершился с ошибкой: %v", err)
	}
}
