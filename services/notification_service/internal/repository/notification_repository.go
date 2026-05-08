package repository

import (
    "gorm.io/gorm"
    "notification-service/internal/entity"
    "notification-service/internal/models"
)

type NotificationRepository struct {
    db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
    return &NotificationRepository{db: db}
}

func entityNotification(data *models.Notification) *entity.NotificationEntity {
    return &entity.NotificationEntity{
        UserID:    data.UserID,
        Channel:   data.Channel,
        Message:   data.Message,
        Status:    data.Status,
        CreatedAt: data.CreatedAt,
        UpdatedAt: data.UpdatedAt,
    }
}

func modelNotification(data *entity.NotificationEntity) *models.Notification {
    return &models.Notification{
        UserID:    data.UserID,
        Channel:   data.Channel,
        Message:   data.Message,
        Status:    data.Status,
        CreatedAt: data.CreatedAt,
        UpdatedAt: data.UpdatedAt,
    }
}

func (n *NotificationRepository) NotificationGet() ([]*entity.NotificationEntity, error) {
    var notificationsModel []*models.Notification

    if err := n.db.Limit(50).Find(&notificationsModel).Error; err != nil {
        return nil, err
    }

    notification := make([]*entity.NotificationEntity, len(notificationsModel))

    for i, v := range notificationsModel {
        notification[i] = entityNotification(v)
    }

    return notification, nil
}

func (n *NotificationRepository) NotificationCreate(data *entity.NotificationEntity) error {
    var notificationModel *models.Notification

    notificationModel = modelNotification(data)

    if err := n.db.Create(notificationModel).Error; err != nil {
        return err
    }
    return nil
}

func (n *NotificationRepository) NotificationDelete(data entity.NotificationEntity) error {

}
