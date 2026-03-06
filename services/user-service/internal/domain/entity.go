package domain

import "time"

type Role string

const (
	RoleCustomer   Role = "customer"
	RoleRestaurant Role = "restaurant"
	RoleAdmin      Role = "admin"
)

type User struct {
	ID           string     `db:"id" json:"id"`
	Name         string     `db:"name" json:"name"`
	Email        string     `db:"email" json:"email"`
	Phone        string     `db:"phone" json:"phone"`
	PasswordHash string     `db:"password_hash" json:"-"`
	Role         Role       `db:"role" json:"role"`
	CreatedAt    *time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updatedAt"`
}
