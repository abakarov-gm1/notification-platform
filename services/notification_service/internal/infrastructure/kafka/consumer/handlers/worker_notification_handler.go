package handlers

import (
	"notification-service/internal/events"
)

type (
	CreateWorkerUseCase interface {
		Send(*events.NotificationEvent) error
	}
)

type WorkerNotificationHandler struct {
	CreateWorkerUseCase CreateWorkerUseCase
}

func NewWorkerNotificationHandler(createWorkerUseCase CreateWorkerUseCase) *WorkerNotificationHandler {
	return &WorkerNotificationHandler{CreateWorkerUseCase: createWorkerUseCase}
}

func (n *WorkerNotificationHandler) CreateWorkerHandler(data *events.NotificationEvent) error {
	err := n.CreateWorkerUseCase.Send(data)
	if err != nil {
		return err
	}
	return nil
}
