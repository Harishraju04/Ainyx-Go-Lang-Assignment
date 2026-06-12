package service

type CreateUserRequest struct {
	Name string `validate:"required,min=3"`
	Dob  string `validate:"required,datetime=2006-01-02"`
}

type User struct {
	Id   int32
	Name string
	Dob  string
	Age  *int32
}

type UpdateUserRequest struct {
	Id   int32   `validate:"required,gt=0"`
	Name *string `validate:"omitempty,min=3"`
	Dob  *string `validate:"omitempty,datetime=2006-01-02"`
}
