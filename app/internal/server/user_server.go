package server

import (
	userpb "auth-service/proto"
	"context"
	"log"

	"auth-service/internal/repository"
)

type UserServer struct {
	userpb.UnimplementedUserServiceServer
	repo repository.UserRepository
}

func NewUserServer(repo repository.UserRepository) *UserServer {
	return &UserServer{
		repo: repo,
	}
}

func (s *UserServer) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	user, err := s.repo.Create(req.GetName())
	if err != nil {
		log.Printf("Ошибка при создании пользователя: %v", err)
		return nil, err
	}
	return &userpb.CreateUserResponse{User: user}, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	users, err := s.repo.FindById(req.GetId())
	if err != nil {
		log.Printf("Ошибка при получении пользователя: %v", err)
		return nil, err
	}
	return &userpb.GetUserResponse{Users: users}, nil
}
