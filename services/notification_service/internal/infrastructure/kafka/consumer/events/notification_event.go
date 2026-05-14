package events

import "time"

type User struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	TelegramID string `json:"telegram_id"`
}

type NotificationEvent struct {
	User      User      `json:"user"`
	Channel   string    `json:"channel"`
	Message   string    `json:"message"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
