package repository

import (
	"context"

	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/logger"
	"go.uber.org/zap"
)

func (r *Repository) GetTotalUserCount(ctx context.Context) (int64, error) {
	count, err := r.q.CountUsers(ctx)
	if err != nil {
		logger.Logger.Error("GetTotalUserCount: failed to count users", zap.Error(err))
		return 0, err
	}
	return count, nil
}
