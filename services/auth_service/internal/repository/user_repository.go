package repository

import (
	"auth-service/internal/entity"
	"auth-service/internal/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func modelToEntity(m *models.User) *entity.User {
	return &entity.User{
		ID:       m.ID,
		Email:    m.Email,
		Password: m.Password,
	}
}

func entityToModel(e *entity.User) *models.User {
	return &models.User{
		ID:       e.ID,
		Email:    e.Email,
		Password: e.Password,
	}
}

// GetUsers -.
func (u *UserRepo) GetUsers() ([]*entity.User, error) {
	var usersModels []*models.User
	if err := u.db.Limit(50).Offset(0).Select("id", "email").Find(&usersModels).Error; err != nil {
		return nil, err
	}
	users := make([]*entity.User, len(usersModels))
	for i, v := range usersModels {
		users[i] = modelToEntity(v)
	}
	return users, nil
}

// GetUser -.
func (u *UserRepo) GetUser(email string) (*entity.User, error) {
	var userModel models.User
	if err := u.db.First(&userModel, "email = ?", email).Error; err != nil {
		return nil, err
	}
	user := modelToEntity(&userModel)
	return user, nil
}

// CreateUser -.
func (u *UserRepo) CreateUser(user *entity.User) error {
	userModel := entityToModel(user)
	err := u.db.Create(&userModel).Error
	return err
}

// DeleteUser -.
func (u *UserRepo) DeleteUser(id uint) error {
	err := u.db.Delete(&models.User{}, id).Error
	return err
}

// UpdateUserEmail -.
func (u *UserRepo) UpdateUserEmail(id uint, email string) error {
	return u.db.Model(&models.User{}).
		Where("id = ?", id).
		Update("email", email).Error
}
