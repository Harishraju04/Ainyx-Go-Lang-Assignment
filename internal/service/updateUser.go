package service

import (
	"context"
	"log"

	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/repository"
)

func (svc *Service) UpdateUser(ctx context.Context, req *UpdateUserRequest) (*User, error) {
	dob, err := ParseDate(*req.Dob)
	if err != nil {
		log.Printf("svc UpdateUser: %s", err)
		return nil, err
	}

	res, err := svc.repo.UpdateUser(ctx, &repository.UpdateUserRequest{
		Id:   req.Id,
		Name: req.Name,
		Dob:  &dob,
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
