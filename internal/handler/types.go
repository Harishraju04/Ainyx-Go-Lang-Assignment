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

type PaginationQuery struct {
	Page     int32 `query:"page" validate:"gt=0"`
	PageSize int32 `query:"pageSize" validate:"gt=0,lte=100"`
}

type PaginatedUsersResponse struct {
	Data       []*User `json:"data"`
	Page       int32   `json:"page"`
	PageSize   int32   `json:"pageSize"`
	Total      int64   `json:"total"`
	TotalPages int32   `json:"totalPages"`
}
