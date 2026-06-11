package repository

import (
	"time"
)

type CreateUserRequest struct {
	Name string
	Dob  time.Time
}

type User struct {
	Id   int32
	Name string
	Dob  time.Time
}

type UpdateUserRequest struct {
	Id   int32
	Name *string
	Dob  *time.Time
}
