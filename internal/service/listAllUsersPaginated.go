package service

import (
	"context"
	"math"

	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/logger"
	"go.uber.org/zap"
)

func (svc *Service) ListAllUsersPaginated(ctx context.Context, page int32, pageSize int32) (*PaginatedUsersResponse, error) {
	offset := (page - 1) * pageSize

	total, err := svc.repo.GetTotalUserCount(ctx)
	if err != nil {
		logger.Logger.Error("ListAllUsersPaginated: failed to get total count", zap.Error(err))
		return nil, err
	}

	res, err := svc.repo.ListAllUsersPaginated(ctx, offset, pageSize)
	if err != nil {
		logger.Logger.Error("ListAllUsersPaginated: failed to list users", zap.Error(err))
		return nil, err
	}

	var users []*User
	for _, user := range res {
		users = append(users, &User{
			Id:   user.Id,
			Name: user.Name,
			Dob:  user.Dob.Format("2006-01-02"),
			Age:  CalculateAge(user.Dob),
		})
	}

	totalPages := int32(math.Ceil(float64(total) / float64(pageSize)))

	return &PaginatedUsersResponse{
		Data:       users,
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: totalPages,
	}, nil
}
