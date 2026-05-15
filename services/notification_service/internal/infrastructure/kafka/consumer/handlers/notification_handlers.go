package handlers

import (
	"notification-service/internal/entity"
	"notification-service/internal/events"
)

type (
	NotificationUseCase interface {
		NotifCreate(newNotif *entity.NotificationEntity) error
	}

	CreateWorkersUseCase interface {
		Send(*events.NotificationEvent) error
	}
)

type NotificationHandler struct {
	NotifUseCase         NotificationUseCase
	CreateWorkersUseCase CreateWorkersUseCase
}

func NewNotificationHandler(notifUseCase NotificationUseCase, createWorkersUseCase CreateWorkersUseCase) *NotificationHandler {
	return &NotificationHandler{NotifUseCase: notifUseCase, CreateWorkersUseCase: createWorkersUseCase}
}

func (n *NotificationHandler) CreateHandler(data *events.NotificationEvent) error {
	newNotif := entity.NotificationEntity{
		UserID:    data.User.Id,
		Channel:   data.Channel,
		Message:   data.Message,
		Status:    data.Status,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	err := n.NotifUseCase.NotifCreate(&newNotif)
	if err != nil {
		return err
	}
	return nil
}

func (n *NotificationHandler) CreateWorkersHandler(data *events.NotificationEvent) error {
	err := n.CreateWorkersUseCase.Send(data)
	if err != nil {
		return err
	}
	return nil
}
