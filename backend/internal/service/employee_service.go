package service

import (
	"erp-backend/internal/model"
	"erp-backend/internal/repository"
)

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) GetAllDepartments() ([]model.Department, error) {
	return s.repo.FindAllDepartments()
}

func (s *employeeService) GetDepartmentByID(id uint) (*model.Department, error) {
	dep, err := s.repo.FindDepartmentByID(id)
	if err != nil {
		return nil, model.ErrNotFound
	}
	return dep, nil
}

func (s *employeeService) CreateDepartment(req model.CreateDepartmentRequest) (*model.Department, error) {
	dep := &model.Department{Name: req.Name, Description: req.Description}
	return dep, s.repo.CreateDepartment(dep)
}

func (s *employeeService) UpdateDepartment(id uint, req model.CreateDepartmentRequest) (*model.Department, error) {
	dep, err := s.repo.FindDepartmentByID(id)
	if err != nil {
		return nil, model.ErrNotFound
	}
	dep.Name = req.Name
	dep.Description = req.Description
	return dep, s.repo.UpdateDepartment(dep)
}

func (s *employeeService) DeleteDepartment(id uint) error {
	if _, err := s.repo.FindDepartmentByID(id); err != nil {
		return model.ErrNotFound
	}
	return s.repo.DeleteDepartment(id)
}

func (s *employeeService) GetAllPositions() ([]model.Position, error) {
	return s.repo.FindAllPositions()
}

func (s *employeeService) CreatePosition(req model.CreatePositionRequest) (*model.Position, error) {
	pos := &model.Position{
		Title:        req.Title,
		DepartmentID: req.DepartmentID,
		MinSalary:    req.MinSalary,
		MaxSalary:    req.MaxSalary,
	}
	return pos, s.repo.CreatePosition(pos)
}

func (s *employeeService) GetAllEmployees() ([]model.Employee, error) {
	return s.repo.FindAllEmployees()
}

func (s *employeeService) GetEmployeeByID(id uint) (*model.Employee, error) {
	emp, err := s.repo.FindEmployeeByID(id)
	if err != nil {
		return nil, model.ErrNotFound
	}
	return emp, nil
}

func (s *employeeService) CreateEmployee(req model.CreateEmployeeRequest) (*model.Employee, error) {
	emp := &model.Employee{
		EmployeeCode: req.EmployeeCode,
		Name:         req.Name,
		Email:        req.Email,
		Phone:        req.Phone,
		DepartmentID: req.DepartmentID,
		PositionID:   req.PositionID,
		Salary:       req.Salary,
		HireDate:     req.HireDate,
		Status:       req.Status,
	}
	if emp.Status == "" {
		emp.Status = model.EmployeeActive
	}
	return emp, s.repo.CreateEmployee(emp)
}

func (s *employeeService) UpdateEmployee(id uint, req model.UpdateEmployeeRequest) (*model.Employee, error) {
	emp, err := s.repo.FindEmployeeByID(id)
	if err != nil {
		return nil, model.ErrNotFound
	}
	if req.Name != "" {
		emp.Name = req.Name
	}
	if req.Email != "" {
		emp.Email = req.Email
	}
	if req.Phone != "" {
		emp.Phone = req.Phone
	}
	if req.DepartmentID != 0 {
		emp.DepartmentID = req.DepartmentID
	}
	if req.PositionID != 0 {
		emp.PositionID = req.PositionID
	}
	if req.Salary != 0 {
		emp.Salary = req.Salary
	}
	if req.Status != "" {
		emp.Status = req.Status
	}
	return emp, s.repo.UpdateEmployee(emp)
}

func (s *employeeService) DeleteEmployee(id uint) error {
	if _, err := s.repo.FindEmployeeByID(id); err != nil {
		return model.ErrNotFound
	}
	return s.repo.DeleteEmployee(id)
}
