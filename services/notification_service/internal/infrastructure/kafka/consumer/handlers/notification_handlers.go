package handlers

import (
	"notification-service/internal/entity"
	"notification-service/internal/events"
)

type (
	NotificationUseCase interface {
		NotifyCreate(newNotify *entity.NotificationEntity) error
	}
)

type NotificationHandler struct {
	NotifyUseCase NotificationUseCase
}

func NewNotificationHandler(notifyUseCase NotificationUseCase) *NotificationHandler {
	return &NotificationHandler{NotifyUseCase: notifyUseCase}
}

func (n *NotificationHandler) CreateHandler(data *events.NotificationEvent) error {
	newNotify := entity.NotificationEntity{
		UserID:    data.User.Id,
		Channel:   data.Channel,
		Message:   data.Message,
		Status:    data.Status,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	err := n.NotifyUseCase.NotifyCreate(&newNotify)
	if err != nil {
		return err
	}
	return nil
}
