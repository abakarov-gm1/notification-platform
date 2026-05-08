package use_case

import (
	"auth-service/internal/entity"
)

type userRepo interface {
	GetUsers() ([]*entity.User, error)
	GetUser(email string) (*entity.User, error)
	CreateUser(user *entity.User) error
	DeleteUser(id uint) error
	UpdateUserEmail(id uint, email string) error
}

type UserUseCase struct {
	repo userRepo
}

func NewUserUseCase(r userRepo) *UserUseCase {
	return &UserUseCase{repo: r}
}

func (u *UserUseCase) Get(email string) {
	//user, err := u.repo.GetUser(email)
}
