package service

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/repository"
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/validator"
)

type fakeRepo struct {
	createFunc func(ctx context.Context, req *repository.CreateUserRequest) (*repository.User, error)
	updateFunc func(ctx context.Context, req *repository.UpdateUserRequest) (*repository.User, error)
}

func (f *fakeRepo) CreateUser(ctx context.Context, req *repository.CreateUserRequest) (*repository.User, error) {
	if f.createFunc != nil {
		return f.createFunc(ctx, req)
	}
	return nil, nil
}

func (f *fakeRepo) UpdateUser(ctx context.Context, req *repository.UpdateUserRequest) (*repository.User, error) {
	if f.updateFunc != nil {
		return f.updateFunc(ctx, req)
	}
	return nil, nil
}

func (f *fakeRepo) GetUserByID(ctx context.Context, id int32) (*repository.User, error) {
	return nil, nil
}

func (f *fakeRepo) ListAllUsers(ctx context.Context) ([]*repository.User, error) {
	return nil, nil
}

func (f *fakeRepo) DeleteUser(ctx context.Context, id int32) error {
	return nil
}

func TestMain(m *testing.M) {
	validator.Init()
	os.Exit(m.Run())
}

func TestService_CreateUser_ValidInput(t *testing.T) {
	fake := &fakeRepo{
		createFunc: func(ctx context.Context, req *repository.CreateUserRequest) (*repository.User, error) {
			if req.Name != "Alice" {
				t.Fatalf("expected name Alice, got %q", req.Name)
			}
			if req.Dob.Format("2006-01-02") != "1990-01-01" {
				t.Fatalf("expected dob 1990-01-01, got %s", req.Dob.Format("2006-01-02"))
			}
			return &repository.User{Id: 1, Name: req.Name, Dob: req.Dob}, nil
		},
	}

	svc := &Service{repo: fake}
	out, err := svc.CreateUser(context.Background(), &CreateUserRequest{Name: "Alice", Dob: "1990-01-01"})
	if err != nil {
		t.Fatal(err)
	}
	if out.Id != 1 || out.Name != "Alice" || out.Dob != "1990-01-01" {
		t.Fatalf("unexpected output: %+v", out)
	}
}

func TestService_UpdateUser_NoFields(t *testing.T) {
	svc := &Service{repo: &fakeRepo{}}
	_, err := svc.UpdateUser(context.Background(), &UpdateUserRequest{Id: 1})
	if err == nil {
		t.Fatal("expected error for no update fields")
	}
	if err.Error() != "at least one field must be provided" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestParseDate_Invalid(t *testing.T) {
	_, err := ParseDate("not-a-date")
	if err == nil {
		t.Fatal("expected error for invalid date")
	}
}

func TestCalculateAge(t *testing.T) {
	dob := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	age := CalculateAge(dob)
	if age == nil {
		t.Fatal("expected non-nil age")
	}
	if *age <= 0 {
		t.Fatalf("expected positive age, got %d", *age)
	}
}
