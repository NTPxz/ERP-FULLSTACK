package model

import (
	"time"

	"gorm.io/gorm"
)

type Department struct {
	gorm.Model
	Name        string     `json:"name"                gorm:"not null;unique"`
	Description string     `json:"description"`
	Employees   []Employee `json:"employees,omitempty" gorm:"foreignKey:DepartmentID"`
}

type Position struct {
	gorm.Model
	Title        string     `json:"title"         gorm:"not null"`
	DepartmentID uint       `json:"department_id" gorm:"not null"`
	Department   Department `json:"department,omitempty"`
	MinSalary    float64    `json:"min_salary"`
	MaxSalary    float64    `json:"max_salary"`
}

type EmployeeStatus string

const (
	EmployeeActive   EmployeeStatus = "active"
	EmployeeInactive EmployeeStatus = "inactive"
)

type Employee struct {
	gorm.Model
	EmployeeCode string         `json:"employee_code" gorm:"unique;not null"`
	Name         string         `json:"name"          gorm:"not null"`
	Email        string         `json:"email"         gorm:"unique;not null"`
	Phone        string         `json:"phone"`
	DepartmentID uint           `json:"department_id"`
	Department   Department     `json:"department,omitempty"`
	PositionID   uint           `json:"position_id"`
	Position     Position       `json:"position,omitempty"`
	Salary       float64        `json:"salary"`
	HireDate     time.Time      `json:"hire_date"`
	Status       EmployeeStatus `json:"status"        gorm:"default:active"`
}

type CreateDepartmentRequest struct {
	Name        string `json:"name"        validate:"required"`
	Description string `json:"description"`
}

type CreatePositionRequest struct {
	Title        string  `json:"title"         validate:"required"`
	DepartmentID uint    `json:"department_id" validate:"required"`
	MinSalary    float64 `json:"min_salary"`
	MaxSalary    float64 `json:"max_salary"`
}

type CreateEmployeeRequest struct {
	EmployeeCode string         `json:"employee_code" validate:"required"`
	Name         string         `json:"name"          validate:"required"`
	Email        string         `json:"email"         validate:"required,email"`
	Phone        string         `json:"phone"`
	DepartmentID uint           `json:"department_id" validate:"required"`
	PositionID   uint           `json:"position_id"   validate:"required"`
	Salary       float64        `json:"salary"`
	HireDate     time.Time      `json:"hire_date"`
	Status       EmployeeStatus `json:"status"`
}

type UpdateEmployeeRequest struct {
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	Phone        string         `json:"phone"`
	DepartmentID uint           `json:"department_id"`
	PositionID   uint           `json:"position_id"`
	Salary       float64        `json:"salary"`
	Status       EmployeeStatus `json:"status"`
}
