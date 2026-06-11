package service

import "context"

func (svc *Service) DeleteUser(ctx context.Context, id int32) error {
	err := svc.repo.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
