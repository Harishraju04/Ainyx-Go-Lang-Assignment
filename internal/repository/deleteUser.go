package repository

import (
	"context"

	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/logger"
	"go.uber.org/zap"
)

func (repo *Repository) DeleteUser(ctx context.Context, id int32) error {
	err := repo.q.DeleteUser(ctx, id)
	if err != nil {
		logger.Logger.Error("DeleteUser: database error", zap.Error(err), zap.Int32("id", id))
		return err
	}
	return nil
}
