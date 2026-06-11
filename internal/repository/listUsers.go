package repository

import (
	"context"
	"log"
)

func (repo *Repository) ListAllUsers(ctx context.Context) ([]*User, error) {
	res, err := repo.q.ListAllUsers(ctx)
	if err != nil {
		log.Printf("repo ListAllUsers %s", err)
		return nil, err
	}
	var users []*User
	for _, user := range res {
		users = append(users, &User{
			Id:   user.ID,
			Name: user.Name,
			Dob:  user.Dob.Time,
		})
	}
	return users, nil
}
