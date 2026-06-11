package service

type CreateUserRequest struct {
	Name string
	Dob  string
}

type User struct {
	Id   int32
	Name string
	Dob  string
	Age  *int32
}

type UpdateUserRequest struct {
	Id   int32
	Name *string
	Dob  *string
}
