package authHandler

import (
	"auth-service/internal/entity"
	"encoding/json"
	"net/http"
)

type (
	AuthUseCase interface {
		Register(email, password string) error
		Login(email, password string) *entity.Auth
	}
)

type AuthHandler struct {
	authUseCase AuthUseCase
}

func NewAuthHandler(authUseCase AuthUseCase) *AuthHandler {
	return &AuthHandler{authUseCase: authUseCase}
}

func (a *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	res := a.authUseCase.Login(req.Email, req.Password)
	if res.Err != nil {
		_, _ = w.Write([]byte("Ошибка, при авторизации"))
		return
	}
	_, _ = w.Write([]byte(res.Token))
}

func (a *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	res := a.authUseCase.Register(req.Email, req.Password)

	if res != nil {
		_, _ = w.Write([]byte("Ошибка при регистрации"))
		return
	}
	_, _ = w.Write([]byte("Регстрация прошла успешно!"))

}
