package use_case

import (
	"notification-service/internal/entity"
)

type NotifRepository interface {
	NotificationCreate(data *entity.NotificationEntity) error
	NotificationDelete(data *entity.NotificationEntity) error
}

type NotificationUseCase struct {
	repo NotifRepository
}

func NewNotificationUseCase(r NotifRepository) *NotificationUseCase {
	return &NotificationUseCase{repo: r}
}

func (n *NotificationUseCase) NotifSend() {

}

func (n *NotificationUseCase) NotifCreate(newNotif *entity.NotificationEntity) error {
	if err := n.repo.NotificationCreate(newNotif); err != nil {
		return err
	}
	return nil
}

func (n *NotificationUseCase) NotifDelete(notif *entity.NotificationEntity) error {
	return n.repo.NotificationDelete(notif)
}

func (n *NotificationUseCase) NotifCreateWorkers() {

}
