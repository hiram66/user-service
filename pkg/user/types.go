package user

import "context"

type User struct {
	Name   string `json:"name"`
	Family string `json:"family"`
	Phone  string `json:"phone"`
	Email  string `json:"email,omitempty"`
}

type UserResponse struct {
	User
	Id string `json:"id"`
}

type UserQuery struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type AddStatus struct {
	Id string `json:"id"`
}

type Service interface {
	AddUser(ctx context.Context, user User) (AddStatus, error)
	GetUserById(ctx context.Context, id string) (UserResponse, error)
	GetUser(ctx context.Context, q UserQuery) (UserResponse, error)
	Update(ctx context.Context, id string, user User) (UserResponse, error)
	Delete(ctx context.Context, id string) error
}
