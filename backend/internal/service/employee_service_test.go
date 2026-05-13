package service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"erp-backend/internal/model"
	"erp-backend/internal/service"
)

type mockEmployeeRepo struct {
	departments []model.Department
	positions   []model.Position
	employees   []model.Employee
	findErr     error
	saveErr     error
}

func (m *mockEmployeeRepo) FindAllDepartments() ([]model.Department, error) {
	return m.departments, m.findErr
}
func (m *mockEmployeeRepo) FindDepartmentByID(id uint) (*model.Department, error) {
	if m.findErr != nil {
		return nil, m.findErr
	}
	for i, d := range m.departments {
		if d.ID == id {
			return &m.departments[i], nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}
func (m *mockEmployeeRepo) CreateDepartment(d *model.Department) error { return m.saveErr }
func (m *mockEmployeeRepo) UpdateDepartment(d *model.Department) error { return m.saveErr }
func (m *mockEmployeeRepo) DeleteDepartment(id uint) error             { return m.saveErr }
func (m *mockEmployeeRepo) FindAllPositions() ([]model.Position, error) {
	return m.positions, m.findErr
}
func (m *mockEmployeeRepo) FindPositionByID(id uint) (*model.Position, error) {
	if m.findErr != nil {
		return nil, m.findErr
	}
	for i, p := range m.positions {
		if p.ID == id {
			return &m.positions[i], nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}
func (m *mockEmployeeRepo) CreatePosition(p *model.Position) error { return m.saveErr }
func (m *mockEmployeeRepo) FindAllEmployees() ([]model.Employee, error) {
	return m.employees, m.findErr
}
func (m *mockEmployeeRepo) FindEmployeeByID(id uint) (*model.Employee, error) {
	if m.findErr != nil {
		return nil, m.findErr
	}
	for i, e := range m.employees {
		if e.ID == id {
			return &m.employees[i], nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}
func (m *mockEmployeeRepo) CreateEmployee(e *model.Employee) error { return m.saveErr }
func (m *mockEmployeeRepo) UpdateEmployee(e *model.Employee) error { return m.saveErr }
func (m *mockEmployeeRepo) DeleteEmployee(id uint) error           { return m.saveErr }

func TestEmployeeService_GetAllDepartments(t *testing.T) {
	tests := []struct {
		name      string
		mock      *mockEmployeeRepo
		wantCount int
		wantErr   bool
	}{
		{
			name:      "returns all departments",
			mock:      &mockEmployeeRepo{departments: []model.Department{{Name: "Engineering"}, {Name: "HR"}}},
			wantCount: 2,
		},
		{
			name:    "propagates repo error",
			mock:    &mockEmployeeRepo{findErr: errors.New("db error")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := service.NewEmployeeService(tt.mock)
			got, err := svc.GetAllDepartments()
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Len(t, got, tt.wantCount)
		})
	}
}

func TestEmployeeService_GetDepartmentByID(t *testing.T) {
	dep := model.Department{Name: "Engineering"}
	dep.ID = 1

	tests := []struct {
		name    string
		mock    *mockEmployeeRepo
		id      uint
		wantErr error
	}{
		{
			name: "returns department when found",
			mock: &mockEmployeeRepo{departments: []model.Department{dep}},
			id:   1,
		},
		{
			name:    "returns ErrNotFound when missing",
			mock:    &mockEmployeeRepo{departments: []model.Department{}},
			id:      999,
			wantErr: model.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := service.NewEmployeeService(tt.mock)
			got, err := svc.GetDepartmentByID(tt.id)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, "Engineering", got.Name)
		})
	}
}

func TestEmployeeService_CreateDepartment(t *testing.T) {
	t.Run("creates department successfully", func(t *testing.T) {
		svc := service.NewEmployeeService(&mockEmployeeRepo{})
		dep, err := svc.CreateDepartment(model.CreateDepartmentRequest{Name: "Finance"})
		assert.NoError(t, err)
		assert.Equal(t, "Finance", dep.Name)
	})

	t.Run("propagates save error", func(t *testing.T) {
		svc := service.NewEmployeeService(&mockEmployeeRepo{saveErr: errors.New("constraint")})
		_, err := svc.CreateDepartment(model.CreateDepartmentRequest{Name: "Finance"})
		assert.Error(t, err)
	})
}

func TestEmployeeService_DeleteDepartment(t *testing.T) {
	dep := model.Department{Name: "Old"}
	dep.ID = 1

	t.Run("deletes existing department", func(t *testing.T) {
		svc := service.NewEmployeeService(&mockEmployeeRepo{departments: []model.Department{dep}})
		err := svc.DeleteDepartment(1)
		assert.NoError(t, err)
	})

	t.Run("returns ErrNotFound for missing department", func(t *testing.T) {
		svc := service.NewEmployeeService(&mockEmployeeRepo{})
		err := svc.DeleteDepartment(999)
		assert.ErrorIs(t, err, model.ErrNotFound)
	})
}

func TestEmployeeService_GetAllEmployees(t *testing.T) {
	t.Run("returns all employees", func(t *testing.T) {
		mock := &mockEmployeeRepo{employees: []model.Employee{{Name: "Alice"}, {Name: "Bob"}}}
		svc := service.NewEmployeeService(mock)
		got, err := svc.GetAllEmployees()
		assert.NoError(t, err)
		assert.Len(t, got, 2)
	})
}
