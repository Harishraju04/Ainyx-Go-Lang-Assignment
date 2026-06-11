package handler

type CreateUserRequest struct {
	Name string `json:"name"`
	Dob  string `json:"dob"`
}

type User struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
	Dob  string `json:"dob"`
	Age  *int32 `json:"age,omitempty"`
}

type UpdateUserRequest struct {
	Name *string `json:"name,omitempty"`
	Dob  *string `json:"dob,omitempty"`
}
