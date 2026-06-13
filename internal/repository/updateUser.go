package repository

import (
	"context"
	"errors"

	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/db/sqlc"
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/logger"
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/validator"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func (repo *Repository) UpdateUser(ctx context.Context, req *UpdateUserRequest) (*User, error) {
	if err := validator.Validate.Struct(req); err != nil {
		return nil, err
	}

	res, err := repo.q.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   req.Id,
		Name: ToText(req.Name),
		Dob:  ToDate(req.Dob),
	})

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Logger.Debug("UpdateUser: no rows found", zap.Int32("id", req.Id))
			return nil, pgx.ErrNoRows
		}
		logger.Logger.Error("UpdateUser: database error", zap.Error(err), zap.Int32("id", req.Id))
		return nil, err
	}

	return &User{
		Id:   res.ID,
		Name: res.Name,
		Dob:  res.Dob.Time,
	}, nil
}
