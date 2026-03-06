package domain

import "errors"

var (
	ErrNotFound        = errors.New("user not found")
	ErrAlreadyExists   = errors.New("user already exists")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidRole     = errors.New("invalid role")
	ErrInvalidInput    = errors.New("invalid input")
	ErrUnauthorized    = errors.New("unauthorized")
)
