package repository

import (
	"context"

	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func (repo *Repository) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
	res, err := repo.q.CreateUser(ctx, sqlc.CreateUserParams{
		Name: req.Name,
		Dob: pgtype.Date{
			Time:  req.Dob,
			Valid: true,
		},
	})

	if err != nil {
		return nil, err
	}

	return &User{
		Id:   res.ID,
		Name: res.Name,
		Dob:  res.Dob.Time,
	}, nil
}
