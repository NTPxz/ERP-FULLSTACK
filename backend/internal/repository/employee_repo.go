package repository

import (
	"erp-backend/internal/model"
	"gorm.io/gorm"
)

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) FindAllDepartments() ([]model.Department, error) {
	var deps []model.Department
	return deps, r.db.Find(&deps).Error
}

func (r *employeeRepository) FindDepartmentByID(id uint) (*model.Department, error) {
	var dep model.Department
	if err := r.db.First(&dep, id).Error; err != nil {
		return nil, err
	}
	return &dep, nil
}

func (r *employeeRepository) CreateDepartment(d *model.Department) error {
	return r.db.Create(d).Error
}

func (r *employeeRepository) UpdateDepartment(d *model.Department) error {
	return r.db.Save(d).Error
}

func (r *employeeRepository) DeleteDepartment(id uint) error {
	return r.db.Delete(&model.Department{}, id).Error
}

func (r *employeeRepository) FindAllPositions() ([]model.Position, error) {
	var pos []model.Position
	return pos, r.db.Preload("Department").Find(&pos).Error
}

func (r *employeeRepository) FindPositionByID(id uint) (*model.Position, error) {
	var pos model.Position
	if err := r.db.Preload("Department").First(&pos, id).Error; err != nil {
		return nil, err
	}
	return &pos, nil
}

func (r *employeeRepository) CreatePosition(p *model.Position) error {
	return r.db.Create(p).Error
}

func (r *employeeRepository) FindAllEmployees() ([]model.Employee, error) {
	var emps []model.Employee
	return emps, r.db.Preload("Department").Preload("Position").Find(&emps).Error
}

func (r *employeeRepository) FindEmployeeByID(id uint) (*model.Employee, error) {
	var emp model.Employee
	if err := r.db.Preload("Department").Preload("Position").First(&emp, id).Error; err != nil {
		return nil, err
	}
	return &emp, nil
}

func (r *employeeRepository) CreateEmployee(e *model.Employee) error {
	return r.db.Create(e).Error
}

func (r *employeeRepository) UpdateEmployee(e *model.Employee) error {
	return r.db.Save(e).Error
}

func (r *employeeRepository) DeleteEmployee(id uint) error {
	return r.db.Delete(&model.Employee{}, id).Error
}
