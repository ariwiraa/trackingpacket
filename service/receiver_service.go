package service

import (
	"weeklytest/domain"
	"weeklytest/repository"
)

func AddReceiver(name, phone string) *domain.Receiver {
	receiver := repository.NewReceiver(name, phone)

	repository.AddReceiver(*receiver)

	return receiver
}
