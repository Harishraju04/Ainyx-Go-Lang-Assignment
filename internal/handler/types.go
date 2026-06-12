package handler

type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=3"`
	Dob  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

type User struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
	Dob  string `json:"dob"`
	Age  *int32 `json:"age,omitempty"`
}

type UpdateUserRequest struct {
	Name *string `json:"name,omitempty" validate:"omitempty,min=3"`
	Dob  *string `json:"dob,omitempty"  validate:"omitempty,datetime=2006-01-02"`
}
