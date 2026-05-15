package use_case

import "notification-service/internal/events"

type Producer interface {
	Send()
}

type CreateWorkersUseCase struct {
	producer Producer
}

func NewWorkerUseCase(producer Producer) *CreateWorkersUseCase {
	return &CreateWorkersUseCase{producer: producer}
}

func (p *CreateWorkersUseCase) Send(data *events.NotificationEvent) error {

	return nil
}
