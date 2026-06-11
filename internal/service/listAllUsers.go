package service

import "context"

func (svc *Service) ListAllUsers(ctx context.Context) ([]*User, error) {
	res, err := svc.repo.ListAllUsers(ctx)
	if err != nil {
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
	return users, nil
}
