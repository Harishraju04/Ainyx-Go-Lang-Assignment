package service

import (
	"context"
)

func (svc *Service) GetUserByID(ctx context.Context, id int32) (*User, error) {
	res, err := svc.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &User{
		Id:   res.Id,
		Name: res.Name,
		Dob:  res.Dob.Format("2006-01-02"),
		Age:  CalculateAge(res.Dob),
	}, nil
}
