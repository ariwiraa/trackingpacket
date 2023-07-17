package service

import (
	"weeklytest/domain"
	"weeklytest/repository"
)

func AddSender(name, phone string) *domain.Sender {
	sender := repository.NewSender(name, phone)

	repository.AddSender(*sender)

	return sender
}

func FindSenderById(id string) *domain.Sender {
	return repository.FindBySenderId(id)
}
