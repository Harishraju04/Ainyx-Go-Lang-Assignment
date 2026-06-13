package repository

import (
	"context"
	"errors"

	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/logger"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func (repo *Repository) GetUserByID(ctx context.Context, id int32) (*User, error) {
	res, err := repo.q.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Logger.Debug("GetUserByID: no rows found", zap.Int32("id", id))
			return nil, pgx.ErrNoRows
		}
		logger.Logger.Error("GetUserByID: database error", zap.Error(err), zap.Int32("id", id))
		return nil, err
	}
	return &User{
		Id:   res.ID,
		Name: res.Name,
		Dob:  res.Dob.Time,
	}, nil
}
