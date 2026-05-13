package service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"erp-backend/internal/model"
	"erp-backend/internal/service"
)

type mockUserRepo struct {
	users   []model.User
	findErr error
	saveErr error
}

func (m *mockUserRepo) FindAll() ([]model.User, error) { return m.users, m.findErr }
func (m *mockUserRepo) FindByID(id uint) (*model.User, error) {
	if m.findErr != nil {
		return nil, m.findErr
	}
	for i, u := range m.users {
		if u.ID == id {
			return &m.users[i], nil
		}
	}
	return nil, errors.New("not found")
}
func (m *mockUserRepo) FindByEmail(email string) (*model.User, error) {
	if m.findErr != nil {
		return nil, m.findErr
	}
	for i, u := range m.users {
		if u.Email == email {
			return &m.users[i], nil
		}
	}
	return nil, errors.New("not found")
}
func (m *mockUserRepo) Create(u *model.User) error { return m.saveErr }
func (m *mockUserRepo) Update(u *model.User) error { return m.saveErr }
func (m *mockUserRepo) Delete(id uint) error       { return m.saveErr }

func TestUserService_GetAll(t *testing.T) {
	t.Run("returns all users", func(t *testing.T) {
		mock := &mockUserRepo{users: []model.User{{Name: "Alice"}, {Name: "Bob"}}}
		svc := service.NewUserService(mock)
		got, err := svc.GetAll()
		assert.NoError(t, err)
		assert.Len(t, got, 2)
	})

	t.Run("propagates repo error", func(t *testing.T) {
		mock := &mockUserRepo{findErr: errors.New("db down")}
		svc := service.NewUserService(mock)
		_, err := svc.GetAll()
		assert.Error(t, err)
	})
}

func TestUserService_Login_WrongPassword(t *testing.T) {
	mock := &mockUserRepo{findErr: errors.New("user not found")}
	svc := service.NewUserService(mock)
	_, err := svc.Login(model.LoginRequest{Email: "nobody@example.com", Password: "x"})
	assert.Error(t, err)
}
