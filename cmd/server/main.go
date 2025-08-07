package main

import (
	"log"

	"github.com/Kilril312/users-service/internal/database"
	transportgrpc "github.com/Kilril312/users-service/internal/transport/grpc"
	"github.com/Kilril312/users-service/internal/user"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Ошибка инициализации БД: %v", err)
	}

	repo := user.NewRepository(db)
	svc := user.NewService(repo)

	if err := transportgrpc.RunGRPC(svc); err != nil {
		log.Fatalf("gRPC сервер завершился с ошибкой: %v", err)
	}
}
