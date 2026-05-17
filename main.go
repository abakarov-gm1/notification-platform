package main

import (
	"encoding/json"
	"fmt"
	"time"
)

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

func main() {

	events := []NotificationEvent{
		{
			User: User{
				Id:         1,
				Name:       "Иван Иванов",
				Email:      "ivan@example.com",
				Phone:      "+79991112233",
				TelegramID: "@ivan_dev",
			},
			Channel:   "telegram",
			Message:   "Ваш заказ успешно оформлен!",
			Status:    "sent",
			CreatedAt: time.Now().Add(-10 * time.Minute),
			UpdatedAt: time.Now().Add(-9 * time.Minute),
		},
		{
			User: User{
				Id:         2,
				Name:       "Анна Петрова",
				Email:      "anna@example.com",
				Phone:      "+79994445566",
				TelegramID: "@anna_p",
			},
			Channel:   "email",
			Message:   "Подтвердите ваш адрес электронной почты.",
			Status:    "pending",
			CreatedAt: time.Now().Add(-5 * time.Minute),
			UpdatedAt: time.Now().Add(-5 * time.Minute),
		},
		{
			User: User{
				Id:         3,
				Name:       "Алексей Сидоров",
				Email:      "alex@example.com",
				Phone:      "+79997778899",
				TelegramID: "",
			},
			Channel:   "sms",
			Message:   "Код для входа: 4815",
			Status:    "failed",
			CreatedAt: time.Now().Add(-1 * time.Hour),
			UpdatedAt: time.Now().Add(-55 * time.Minute),
		},
	}

	value, _ := json.Marshal(events[0])

	fmt.Println(value)

	var h NotificationEvent

	json.Unmarshal(value, &h)

	fmt.Println(h, "dsdsdsds")

}
