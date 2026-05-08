package main

import (
	_ "auth-service/docs"
	"auth-service/internal/controller/rest_api"
	"auth-service/internal/database"
	"auth-service/internal/infra/auth_infra"
	"auth-service/internal/infra/jwt"
	"auth-service/internal/models"
	"auth-service/internal/repository"
	"auth-service/internal/use_case"
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"log"
	"net/http"
)

func main() {

	connect, err := database.NewPostgresConnection()
	if err != nil {
		fmt.Println("No Connect Database")
	}
	if err := connect.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Ошибка миграции:", err)
	}
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	userRepo := repository.NewUserRepository(connect)
	hash := auth_infra.NewAuthHashing()
	jwtInfra := jwt.NewJwt("secret")
	AuthUseCase := use_case.NewAuthCase(userRepo, hash, jwtInfra)
	r := rest_api.NewRouter(AuthUseCase)
	_ = http.ListenAndServe(":8081", r)

}

// добавть ошибки
// исправить роутер чтоб обрабатывал ошибки

// добавить сваггер - будет постман пофиг

// в роутах починить кейсы и разобраться с интерфейсами
// протестить
// добавить тесты
// добавить логи
