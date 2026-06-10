package service

import (
	"context"
	"log"
	"time"

	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/repository"
)

func (svc *Service) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	dob, err := time.Parse(
		"2006-01-02",
		req.Dob,
	)
	log.Println("svc", dob)
	if err != nil {
		log.Printf("svc CreateUser: %s", err)
		return nil, err
	}
	res, err := svc.repo.CreateUser(ctx, &repository.CreatUserInput{
		Name: req.Name,
		Dob:  dob,
	})

	if err != nil {
		log.Printf("svc CreateUser %s", err)
		return nil, err
	}

	return &CreateUserResponse{
		Id:   res.Id,
		Name: res.Name,
		Dob:  res.Dob.Format("2006-01-02"),
	}, nil
}
