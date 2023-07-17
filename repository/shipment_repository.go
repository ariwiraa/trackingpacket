package repository

import (
	"errors"
	"strings"
	"weeklytest/domain"
	"weeklytest/helpers"

	"github.com/google/uuid"
)

type ListShipment struct {
	Shipments []domain.Shipment `json:"shipments"`
}

func NewShipment(price float64, isReceived bool, service domain.Service, packet domain.Packet, checkpoint []domain.Location) *domain.Shipment {
	id := "shipment-" + uuid.New().String()
	return &domain.Shipment{Id: id, Price: price, IsReceived: isReceived, Service: service, Packet: packet, CheckPoint: checkpoint}
}

func AddShipment(shipment domain.Shipment) {
	listShipment := ListShipment{}

	path := "/home/ariwiraa/go/src/phincon/week2/weeklytest/files/shipment.json"
	helpers.OpenFile(path, &listShipment)

	listShipment.Shipments = append(listShipment.Shipments, shipment)

	helpers.WriteToFileJson(path, listShipment)
}

func FindShipmentById(id string) (*domain.Shipment, error) {
	listShipment := ListShipment{}

	path := "/home/ariwiraa/go/src/phincon/week2/weeklytest/files/shipment.json"
	helpers.OpenFile(path, &listShipment)

	for i, shipment := range listShipment.Shipments {
		if strings.EqualFold(shipment.Id, id) {
			return &listShipment.Shipments[i], nil
		}
	}
	return nil, errors.New("id shipment tidak ditemukan")
}

func FindAllShipments() []domain.Shipment {
	listShipment := ListShipment{}

	path := "/home/ariwiraa/go/src/phincon/week2/weeklytest/files/shipment.json"
	helpers.OpenFile(path, &listShipment)

	return listShipment.Shipments
}

func UpdateShipment(id string, isReceived bool, location domain.Location) {
	listShipment := ListShipment{}

	path := "/home/ariwiraa/go/src/phincon/week2/weeklytest/files/shipment.json"
	helpers.OpenFile(path, &listShipment)

	for i, shipment := range listShipment.Shipments {
		if strings.EqualFold(shipment.Id, id) {
			listShipment.Shipments[i].IsReceived = isReceived
			listShipment.Shipments[i].CheckPoint = append(listShipment.Shipments[i].CheckPoint, location)
		}
	}

	helpers.WriteToFileJson(path, listShipment)

}
