package model

import "gorm.io/gorm"

type UserRole string

const (
	RoleAdmin     UserRole = "admin"
	RoleHR        UserRole = "hr"
	RoleInventory UserRole = "inventory"
	RoleSales     UserRole = "sales"
	RolePurchase  UserRole = "purchase"
)

type User struct {
	gorm.Model
	Name     string   `json:"name"  gorm:"not null"`
	Email    string   `json:"email" gorm:"unique;not null"`
	Password string   `json:"-"`
	Role     UserRole `json:"role"  gorm:"not null;default:'hr'"`
}

type RegisterRequest struct {
	Name     string `json:"name"     validate:"required"`
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AdminCreateUserRequest struct {
	Name     string   `json:"name"     validate:"required"`
	Email    string   `json:"email"    validate:"required,email"`
	Password string   `json:"password" validate:"required,min=6"`
	Role     UserRole `json:"role"     validate:"required"`
}

type AdminUpdateUserRequest struct {
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Role  UserRole `json:"role"`
}
