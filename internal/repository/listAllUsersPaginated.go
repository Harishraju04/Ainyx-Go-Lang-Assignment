package repository

import (
	"context"

	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/db/sqlc"
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/logger"
	"go.uber.org/zap"
)

func (r *Repository) ListAllUsersPaginated(ctx context.Context, offset int32, limit int32) ([]*User, error) {
	params := sqlc.GetUsersPaginatedParams{
		Offset: offset,
		Limit:  limit,
	}

	rows, err := r.q.GetUsersPaginated(ctx, params)
	if err != nil {
		logger.Logger.Error("ListAllUsersPaginated: failed to query users", zap.Error(err), zap.Int32("offset", offset), zap.Int32("limit", limit))
		return nil, err
	}

	var users []*User
	for _, row := range rows {
		users = append(users, &User{
			Id:   row.ID,
			Name: row.Name,
			Dob:  row.Dob.Time,
		})
	}
	return users, nil
}
