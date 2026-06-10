package service

import "github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/repository"

type Service struct {
	repo *repository.Repository
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		repo: r,
	}
}
