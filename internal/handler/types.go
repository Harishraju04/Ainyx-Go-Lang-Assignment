package handler

type CreateUserRequest struct {
	Name string `json:"name"`
	Dob  string `json:"dob"`
}

type CreateUserResponse struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
	Dob  string `json:"dob"`
}
