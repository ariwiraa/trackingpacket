package repository

import (
	"strings"
	"weeklytest/domain"
	"weeklytest/helpers"

	"github.com/google/uuid"
)

// var senders []domain.Sender

type ListSender struct {
	Senders []domain.Sender `json:"senders"`
}

func NewSender(name, phone string) *domain.Sender {
	id := "sender-" + uuid.New().String()
	return &domain.Sender{Id: id, Name: name, Phone: phone}
}

func AddSender(sender domain.Sender) {
	listSenders := ListSender{}

	path := "/home/ariwiraa/go/src/phincon/week2/weeklytest/files/sender.json"
	helpers.OpenFile(path, &listSenders)

	listSenders.Senders = append(listSenders.Senders, sender)

	helpers.WriteToFileJson(path, listSenders)
}

func FindBySenderId(id string) *domain.Sender {
	// var sender domain.Sender
	listSenders := ListSender{}

	path := "/home/ariwiraa/go/src/phincon/week2/weeklytest/files/sender.json"
	helpers.OpenFile(path, &listSenders)

	for _, send := range listSenders.Senders {
		if strings.EqualFold(send.Id, id) {
			return &send
		}
	}

	return nil
}
