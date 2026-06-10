package repository

import (
	"context"

	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func (repo *Repository) CreateUser(ctx context.Context, input *CreatUserInput) (*CreateUserOutput, error) {
	res, err := repo.q.CreateUser(ctx, sqlc.CreateUserParams{
		Name: input.Name,
		Dob: pgtype.Date{
			Time:  input.Dob,
			Valid: true,
		},
	})

	if err != nil {
		return nil, err
	}

	return &CreateUserOutput{
		Id:   res.ID,
		Name: res.Name,
		Dob:  res.Dob.Time,
	}, nil
}
