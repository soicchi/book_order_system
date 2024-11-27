package users

import (
	"github.com/soicchi/book_order_system/internal/usecase/users"
)

type UserResponse struct {
	User *users.RetrieveOutput `json:"user"`
}

func NewUserResponse(user *users.RetrieveOutput) *UserResponse {
	return &UserResponse{
		User: user,
	}
}

type UsersResponse struct {
	Users []*users.ListOutput `json:"users"`
}

func NewUsersResponse(users []*users.ListOutput) *UsersResponse {
	return &UsersResponse{
		Users: users,
	}
}
