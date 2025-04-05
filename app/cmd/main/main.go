package main

import (
	userpb "auth-service/proto"
	"database/sql"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	postgresRepository "auth-service/internal/infrastructure/postgres/repository"
	"auth-service/internal/repository"
	"auth-service/internal/server"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL не установлен")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("[main] database connection failed: %v", err)
	}
	log.Println("[main] database connection opened")
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("ping error: %v", err)
	}

	var userRepo repository.UserRepository
	userRepo = postgresRepository.NewPostgresUserRepository(db)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("[main] failed to listen: %v", err)
	}
	log.Println("[main] gRPC server listening on :50051")

	grpcServer := grpc.NewServer()
	userServer := server.NewUserServer(userRepo)
	userpb.RegisterUserServiceServer(grpcServer, userServer)

	log.Println("gRPC сервер запущен на порту :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("ошибка запуска сервера: %v", err)
	}
}
