package handlers

import (
	"notification-service/internal/entity"
	"notification-service/internal/infrastructure/kafka/consumer/events"
)

type NotificationUseCase interface {
	NotifCreate(newNotif *entity.NotificationEntity) error
}

type NotificationHandler struct {
	notifUseCase NotificationUseCase
}

func NewNotificationHandler(notifUseCase NotificationUseCase) *NotificationHandler {
	return &NotificationHandler{notifUseCase: notifUseCase}
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

	err := n.notifUseCase.NotifCreate(&newNotif)
	if err != nil {
		return err
	}
	return nil
}

func (n *NotificationHandler) CreateWorkersHandler(data *events.NotificationEvent) error {
	return nil
	// тут создаю нужные разделения и отправляю в консюмер
}
