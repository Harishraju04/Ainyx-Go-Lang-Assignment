package repository

import (
	"context"
	"log"

	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/db/sqlc"
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/validator"
)

func (repo *Repository) UpdateUser(ctx context.Context, req *UpdateUserRequest) (*User, error) {
	if err := validator.Validate.Struct(req); err != nil {
		return nil, err
	}

	res, err := repo.q.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   req.Id,
		Name: ToText(req.Name),
		Dob:  ToDate(req.Dob),
	})

	if err != nil {
		log.Printf("repo UpdateUser: %s", err)
		return nil, err
	}

	return &User{
		Id:   res.ID,
		Name: res.Name,
		Dob:  res.Dob.Time,
	}, nil
}
