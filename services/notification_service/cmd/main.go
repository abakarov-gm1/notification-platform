package main

import (
	"fmt"
	"notification-service/internal/database"
	"notification-service/internal/infrastructure/kafka/consumer"
	"notification-service/internal/infrastructure/kafka/consumer/handlers"
	"notification-service/internal/infrastructure/kafka/producer"
	"notification-service/internal/repository"
	"notification-service/internal/use_case"
)

func main() {

	db, err := database.NewConnection()
	if err != nil {
		fmt.Errorf("error conncect database %e", err)
	}

	repo := repository.NewNotificationRepository(db)
	newProducer := producer.NewProducer()

	newWorkerUseCase := use_case.NewWorkerUseCase(newProducer)
	newNotifyUseCase := use_case.NewNotificationUseCase(repo)

	newWorkerHandler := handlers.NewWorkerNotificationHandler(newWorkerUseCase)
	newNotifyHandler := handlers.NewNotificationHandler(newNotifyUseCase)

	newConsumer := consumer.NewConsumer(newNotifyHandler, newWorkerHandler)

	newConsumer.Start()

}
