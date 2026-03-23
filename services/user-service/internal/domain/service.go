package domain

import "context"

type UserService interface {
	Register(ctx context.Context, req RegisterInput) (*User, string, error)
	Login(ctx context.Context, req LoginInput) (*User, string, error)
	GetUser(ctx context.Context, id string) (*User, error)
	UpdateUser(ctx context.Context, req UpdateInput) (*User, error)
}

type RegisterInput struct {
	Name     string
	Email    string
	Password string
}

type UpdateInput struct {
	ID    string
	Name  string
	Phone string
	Role  string
}

type LoginInput struct {
	Email    string
	Password string
}
