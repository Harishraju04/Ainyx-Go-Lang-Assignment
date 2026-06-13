package repository

import (
	"context"

	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/logger"
	"go.uber.org/zap"
)

func (repo *Repository) ListAllUsers(ctx context.Context) ([]*User, error) {
	res, err := repo.q.ListAllUsers(ctx)
	if err != nil {
		logger.Logger.Error("ListAllUsers: database error", zap.Error(err))
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
