package repository

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
)

func (repo *Repository) GetUserByID(ctx context.Context, id int32) (*User, error) {
	res, err := repo.q.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pgx.ErrNoRows
		}
		log.Printf("Repo GetUserByID: %s", err)
		return nil, err
	}
	return &User{
		Id:   res.ID,
		Name: res.Name,
		Dob:  res.Dob.Time,
	}, nil
}
