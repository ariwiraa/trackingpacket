package service

import (
	"strings"
	"weeklytest/domain"
	"weeklytest/repository"
)

func AddPacket(request domain.RequestPacket) *domain.Packet {
	sender := repository.FindBySenderId(request.SenderId)
	receiver := repository.FindByReceiverId(request.ReceiverId)
	destination := repository.FindLocationByName(request.Destination)
	start := repository.FindLocationByName(request.StartLocation)

	packet := repository.NewPacket(request.Weight, *sender, *receiver, *start, *destination)

	repository.AddPacket(*packet)
	return packet
}

func ListPacketReceived() []domain.PacketDetails {
	var results []domain.PacketDetails

	shipments := repository.FindAllShipments()

	for _, shipment := range shipments {
		if shipment.IsReceived {
			packet := repository.NewPacketDetails(shipment.Packet, shipment.IsReceived)
			results = append(results, *packet)
		}
	}

	return results
}

func GetAllPacketsByLocationName(name string) []domain.PacketDetails {
	var results []domain.PacketDetails

	shipments := repository.FindAllShipments()

	for _, shipment := range shipments {
		for _, location := range shipment.CheckPoint {
			if strings.EqualFold(location.Name, name) {
				packet := repository.NewPacketDetails(shipment.Packet, shipment.IsReceived)
				results = append(results, *packet)
			}
		}
	}

	return results
}
