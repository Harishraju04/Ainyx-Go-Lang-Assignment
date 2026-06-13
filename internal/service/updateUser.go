package service

import (
	"context"
	"errors"
	"time"

	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/logger"
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/repository"
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/validator"
	"go.uber.org/zap"
)

func (svc *Service) UpdateUser(ctx context.Context, req *UpdateUserRequest) (*User, error) {
	if err := validator.Validate.Struct(req); err != nil {
		return nil, err
	}

	if req.Name == nil && req.Dob == nil {
		return nil, errors.New("at least one field must be provided")
	}

	var dob *time.Time
	if req.Dob != nil {
		parsedDob, err := ParseDate(*req.Dob)
		if err != nil {
			logger.Logger.Error("UpdateUser: invalid date format", zap.Error(err))
			return nil, err
		}
		dob = &parsedDob
	}

	res, err := svc.repo.UpdateUser(ctx, &repository.UpdateUserRequest{
		Id:   req.Id,
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
