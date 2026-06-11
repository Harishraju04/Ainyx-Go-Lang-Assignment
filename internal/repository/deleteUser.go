package repository

import (
	"context"
	"log"
)

func (repo *Repository) DeleteUser(ctx context.Context, id int32) error {
	err := repo.q.DeleteUser(ctx, id)
	if err != nil {
		log.Printf("repo DeleteUser %s", err)
		return err
	}
	return nil
}
