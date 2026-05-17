package use_case

import (
	"encoding/json"
	"fmt"
	"notification-service/internal/events"
)

type Producer interface {
	Send(topic string, value []byte) error
	Close()
}

type CreateWorkersUseCase struct {
	producer Producer
}

func NewWorkerUseCase(producer Producer) *CreateWorkersUseCase {
	return &CreateWorkersUseCase{producer: producer}
}

func (p *CreateWorkersUseCase) Send(data *events.NotificationEvent) error {

	topics := map[string][]*events.NotificationEvent{
		"email":    make([]*events.NotificationEvent, 0),
		"telegram": make([]*events.NotificationEvent, 0),
		"sms":      make([]*events.NotificationEvent, 0),
	}

	user := data.User

	if user.Email != "" {
		topics["email"] = append(topics["email"], data)
	}

	if user.TelegramID != "" {
		topics["telegram"] = append(topics["telegram"], data)
	}

	if user.Phone != "" {
		topics["sms"] = append(topics["sms"], data)
	}

	for channel, eventsList := range topics {
		if len(eventsList) == 0 {
			continue // Пропускаем пустые каналы
		}
		value, err := json.Marshal(eventsList)
		if err != nil {
			fmt.Errorf("error from marhal %e", err)
		}
		err = p.producer.Send(channel, value)
		if err != nil {
			fmt.Errorf("error from Producer Send %e", err)
		}
	}
	return nil
}
