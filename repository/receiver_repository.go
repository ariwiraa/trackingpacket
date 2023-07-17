package repository

import (
	"strings"
	"weeklytest/domain"
	"weeklytest/helpers"

	"github.com/google/uuid"
)

type ListReceiver struct {
	Receivers []domain.Receiver `json:"receivers"`
}

func NewReceiver(name, phone string) *domain.Receiver {
	id := "receiver-" + uuid.New().String()
	return &domain.Receiver{Id: id, Name: name, Phone: phone}
}

func AddReceiver(receiver domain.Receiver) {
	listReceiver := ListReceiver{}

	path := "/home/ariwiraa/go/src/phincon/week2/weeklytest/files/receiver.json"
	helpers.OpenFile(path, &listReceiver)

	listReceiver.Receivers = append(listReceiver.Receivers, receiver)

	helpers.WriteToFileJson(path, listReceiver)
}

func FindByReceiverId(id string) *domain.Receiver {
	listReceivers := ListReceiver{}

	path := "/home/ariwiraa/go/src/phincon/week2/weeklytest/files/receiver.json"
	helpers.OpenFile(path, &listReceivers)

	for _, receiver := range listReceivers.Receivers {
		if strings.EqualFold(receiver.Id, id) {
			return &receiver
		}
	}

	return nil
}
