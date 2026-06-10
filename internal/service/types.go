package service

type CreateUserRequest struct {
	Name string
	Dob  string
}

type CreateUserResponse struct {
	Id   int32
	Name string
	Dob  string
}
