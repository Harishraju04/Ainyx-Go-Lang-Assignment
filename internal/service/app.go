package service

import (
	"context"

	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/repository"
)

type repo interface {
	CreateUser(ctx context.Context, req *repository.CreateUserRequest) (*repository.User, error)
	UpdateUser(ctx context.Context, req *repository.UpdateUserRequest) (*repository.User, error)
	GetUserByID(ctx context.Context, id int32) (*repository.User, error)
	ListAllUsers(ctx context.Context) ([]*repository.User, error)
	DeleteUser(ctx context.Context, id int32) error
}

type Service struct {
	repo repo
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		repo: r,
	}
}
