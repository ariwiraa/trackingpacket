package services

import (
	"fmt"
	"strings"
)

var deliveries = []Delivery{}

type Delivery struct {
	Id         string
	Price      float64
	IsReceived bool
	Service    Service
	Packet     Packet
	CheckPoint []Location
}

func NewDelivery(id string, price float64, isReceived bool, service Service, packet Packet, checkPoint []Location) *Delivery {
	return &Delivery{Id: id, Price: price, IsReceived: isReceived, Service: service, Packet: packet, CheckPoint: checkPoint}
}

func AddDelivery(packet *Packet, service *Service) {
	id := len(deliveries) + 1
	idPacket := fmt.Sprintf("del-%d", id)
	price := packet.Weight * service.PriceKg

	d := NewDelivery(idPacket, price, false, *service, *packet, []Location{})
	deliveries = append(deliveries, *d)
}

func ListPacketReceived() []PacketDetails {
	var results []PacketDetails

	for _, v := range deliveries {
		if v.IsReceived {
			packet := NewPacketDetails(v.Packet, v.IsReceived)
			results = append(results, *packet)
		}
	}

	return results
}

func UpdateCheckPoint(deliveryId, locationId string) *Delivery {
	shipment := FindDeliveryById(deliveryId)
	locationObj := FindLocationById(locationId)

	if shipment == nil || locationObj == nil {
		return nil
	}

	shipment.CheckPoint = append(shipment.CheckPoint, *locationObj)
	// fmt.Println(shipment.CheckPoint)

	if len(shipment.CheckPoint) > 0 {
		check := shipment.CheckPoint[len(shipment.CheckPoint)-1]
		if check.Id == shipment.Packet.Destination.Id {
			shipment.IsReceived = true
		}
	}

	return shipment
}

func GetAllPacketsByLocationName(locationName string) []Packet {

	var results []Packet

	for _, delivery := range deliveries {
		for _, loc := range delivery.CheckPoint {
			if strings.EqualFold(loc.Name, locationName) {
				results = append(results, delivery.Packet)
			}
		}
	}

	return results
}

func FindDeliveryById(deliveryId string) *Delivery {
	for i, del := range deliveries {
		if strings.EqualFold(del.Id, deliveryId) {
			return &deliveries[i]
		}
	}

	return nil
}
