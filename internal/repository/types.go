package repository

import (
	"time"
)

type CreatUserInput struct {
	Name string
	Dob  time.Time
}

type CreateUserOutput struct {
	Id   int32
	Name string
	Dob  time.Time
}
