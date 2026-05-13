package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"erp-backend/internal/model"
	"erp-backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAll() ([]model.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetByID(id uint) (*model.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, model.ErrNotFound
	}
	return user, nil
}

func (s *userService) Register(req model.RegisterRequest) (*model.User, error) {
	existing, _ := s.repo.FindByEmail(req.Email)
	if existing != nil {
		return nil, model.ErrConflict
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashed),
		Role:     model.RoleHR, // default role for self-registration
	}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Login(req model.LoginRequest) (string, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    string(user.Role),
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signed, nil
}

func (s *userService) Update(id uint, req model.UpdateUserRequest) (*model.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, model.ErrNotFound
	}
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if err := s.repo.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Delete(id uint) error {
	if _, err := s.repo.FindByID(id); err != nil {
		return model.ErrNotFound
	}
	return s.repo.Delete(id)
}

func (s *userService) AdminCreate(req model.AdminCreateUserRequest) (*model.User, error) {
	existing, _ := s.repo.FindByEmail(req.Email)
	if existing != nil {
		return nil, model.ErrConflict
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashed),
		Role:     req.Role,
	}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) AdminUpdate(id uint, req model.AdminUpdateUserRequest) (*model.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, model.ErrNotFound
	}
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Role != "" {
		user.Role = req.Role
	}
	if err := s.repo.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}
