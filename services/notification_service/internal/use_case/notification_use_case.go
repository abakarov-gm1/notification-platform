package use_case

import (
	"notification-service/internal/entity"
)

type NotifyRepository interface {
	NotificationCreate(data *entity.NotificationEntity) error
	NotificationDelete(data *entity.NotificationEntity) error
}

type NotificationUseCase struct {
	repo NotifyRepository
}

func NewNotificationUseCase(r NotifyRepository) *NotificationUseCase {
	return &NotificationUseCase{repo: r}
}

func (n *NotificationUseCase) NotifySend() {

}

func (n *NotificationUseCase) NotifyCreate(newNotify *entity.NotificationEntity) error {
	if err := n.repo.NotificationCreate(newNotify); err != nil {
		return err
	}
	return nil
}

func (n *NotificationUseCase) NotifyDelete(notify *entity.NotificationEntity) error {
	return n.repo.NotificationDelete(notify)
}

func (n *NotificationUseCase) NotifyCreateWorkers() {

}
