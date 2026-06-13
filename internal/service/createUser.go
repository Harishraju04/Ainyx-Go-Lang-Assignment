package service

import (
	"context"

	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/logger"
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/repository"
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/validator"
	"go.uber.org/zap"
)

func (svc *Service) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
	if err := validator.Validate.Struct(req); err != nil {
		return nil, err
	}

	dob, err := ParseDate(req.Dob)
	if err != nil {
		logger.Logger.Error("CreateUser: invalid date format", zap.Error(err))
		return nil, err
	}
	res, err := svc.repo.CreateUser(ctx, &repository.CreateUserRequest{
		Name: req.Name,
		Dob:  dob,
	})

	if err != nil {
		return nil, err
	}

	return &User{
		Id:   res.Id,
		Name: res.Name,
		Dob:  res.Dob.Format("2006-01-02"),
	}, nil
}
