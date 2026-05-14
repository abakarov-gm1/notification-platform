package rest_api

import (
	"auth-service/internal/controller/rest_api/authHandler"
	"auth-service/internal/entity"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type (
	AuthUseCase interface {
		Register(email, password string) error
		Login(email, password string) *entity.Auth
	}
)

func NewRouter(a AuthUseCase) http.Handler {

	auth := authHandler.NewAuthHandler(a)

	r := chi.NewRouter()

	r.Get("/", helloHandler)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", auth.LoginHandler)
		r.Post("/register", auth.RegisterHandler)
	})

	return r
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello worlds from go"))
}
