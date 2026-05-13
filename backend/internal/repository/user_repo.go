package repository

import (
	"erp-backend/internal/model"
	"gorm.io/gorm"
)

// Struct ที่ implement interface
type userRepository struct {
	db *gorm.DB
}

// Constructor — pattern มาตรฐานใน Go
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindAll() ([]model.User, error) {
	var users []model.User
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *userRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}
