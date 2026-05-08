package use_case

import (
	"auth-service/internal/entity"
	"errors"
)

type (
	authRepo interface {
		GetUsers() ([]*entity.User, error)
		GetUser(email string) (*entity.User, error)
		CreateUser(user *entity.User) error
	}

	hash interface {
		Hashing(s string) ([]byte, error)
		CheckPassword(hashedPassword, userEnteredPassword string) error
	}

	tokenIssuer interface {
		GenerateToken(userId uint) (string, error)
		DecodeToken(tokenString string) (uint, error)
	}
)

type AuthUse struct {
	repo    authRepo
	hashing hash
	token   tokenIssuer
}

func NewAuthCase(r authRepo, hashing hash, token tokenIssuer) *AuthUse {
	return &AuthUse{
		repo:    r,
		hashing: hashing,
		token:   token,
	}
}

func (a *AuthUse) Register(email, password string) error {
	hashPassword, err := a.hashing.Hashing(password)
	if err != nil {
		return err
	}
	newUser := &entity.User{Email: email, Password: string(hashPassword)}
	if err := a.repo.CreateUser(newUser); err != nil {
		return err
	}
	return nil
}

func (a *AuthUse) Login(email, password string) *entity.Auth {
	user, err := a.repo.GetUser(email)

	resp := &entity.Auth{Token: "", Err: nil}

	if err != nil {
		resp.Err = err
		return resp
	}
	if user == nil {
		resp.Err = errors.New("ошибка пользователь не найден")
		return resp
	}
	if err := a.hashing.CheckPassword(user.Password, password); err != nil {
		resp.Err = err
		return resp
	}

	token, er := a.token.GenerateToken(user.ID)
	if er != nil {
		resp.Err = er
		return resp
	}
	resp.Token = token
	return resp
}
