package entity

import "time"

type NotificationEntity struct {
	UserID    uint      `json:"user_id"`
	Channel   string    `json:"channel"`
	Message   string    `json:"message"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
