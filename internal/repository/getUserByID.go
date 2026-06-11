package repository

import (
	"context"
	"log"
)

func (repo *Repository) GetUserByID(ctx context.Context, id int32) (*User, error) {
	res, err := repo.q.GetUserByID(ctx, id)
	if err != nil {
		log.Printf("Repo GetUserByID: %s", err)
		return nil, err
	}
	return &User{
		Id:   res.ID,
		Name: res.Name,
		Dob:  res.Dob.Time,
	}, nil
}
