package domain

import "context"

type UserService interface {
	Register(ctx context.Context, req RegisterInput) (*User, error)
	Login(ctx context.Context, req LoginInput) (*User, error)
	GetUser(ctx context.Context, id string) (*User, error)
	UpdateUser(ctx context.Context, req UpdateInput) (*User, error)
}

type RegisterInput struct {
	Name     string
	Email    string
	Password string
	Phone    string
	Role     Role
}

type UpdateInput struct {
	Id    string
	Name  string
	Phone string
}

type LoginInput struct {
	Email    string
	Password string
}
