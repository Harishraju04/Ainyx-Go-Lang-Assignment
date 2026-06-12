package repository

import (
	"time"
)

type CreateUserRequest struct {
	Name string    `validate:"required,min=3"`
	Dob  time.Time `validate:"required"`
}

type User struct {
	Id   int32
	Name string
	Dob  time.Time
}

type UpdateUserRequest struct {
	Id   int32      `validate:"required,gt=0"`
	Name *string    `validate:"omitempty,min=3"`
	Dob  *time.Time `validate:"omitempty"`
}
