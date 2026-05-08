package models

import (
    "github.com/google/uuid"
    "time"
)

type Notification struct {
    // gorm:"type:uuid;primaryKey" указывает Postgres использовать тип uuid
    // default:gen_random_uuid() заставляет БД генерировать ID автоматически, если он пуст
    ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
    UserID    uint      `gorm:"not null;index"` // индекс для быстрого поиска уведомлений юзера
    Channel   string    `gorm:"type:varchar(50);not null"`
    Message   string    `gorm:"type:text;not null"`
    Status    string    `gorm:"type:varchar(20);default:'pending'"`
    CreatedAt time.Time `gorm:"autoCreateTime"` // GORM заполнит сам при создании
    UpdatedAt time.Time `gorm:"autoUpdateTime"` // GORM обновит сам при сохранении
}
