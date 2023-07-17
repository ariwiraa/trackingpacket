package service

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
	"weeklytest/domain"
	"weeklytest/repository"
)

func AddShipment(request domain.RequestShipment) (*domain.Shipment, error) {
	packet, _ := repository.FindPacketById(request.PacketId)
	service, _ := repository.FindServiceByName(request.ServiceName)
	price := packet.Weight * service.PriceKg

	shipment := repository.NewShipment(price, false, *service, *packet, []domain.Location{packet.StartLocation})

	repository.AddShipment(*shipment)

	return shipment, nil
}

func UpdateCheckPoint(shipmentId, locationName string) *domain.Shipment {
	shipment, _ := repository.FindShipmentById(shipmentId)
	location := repository.FindLocationByName(locationName)

	if shipment == nil || location == nil {
		return nil
	}

	shipment.CheckPoint = append(shipment.CheckPoint, *location)

	if len(shipment.CheckPoint) > 0 {
		check := shipment.CheckPoint[len(shipment.CheckPoint)-1]
		if strings.EqualFold(check.Id, shipment.Packet.Destination.Id) {
			shipment.IsReceived = true
		}
	}

	repository.UpdateShipment(shipment.Id, shipment.IsReceived, *location)

	return shipment
}

func GetShipmentById(id string) (*domain.Shipment, error) {
	shipment, err := repository.FindShipmentById(id)
	if err != nil {
		return shipment, errors.New("pengiriman tidak ditemukan")
	}

	return shipment, nil
}

func CreateFileCSV() {
	shipments := repository.FindAllShipments()
	file, err := os.Create("/home/ariwiraa/go/src/phincon/week2/weeklytest/files/shipment.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	header := []string{"Id Shipment", "Service", "Weight", "Price", "Sender Name", "Received Name", "Destination", "Is Received"}
	err = writer.Write(header)
	if err != nil {
		log.Fatal(err)
	}
	data := make([][]string, 0)
	for _, shipment := range shipments {
		data = append(data, []string{shipment.Id, shipment.Service.Name, strconv.FormatFloat(shipment.Packet.Weight, 'f', 1, 64), strconv.FormatFloat(shipment.Price, 'f', 1, 64), shipment.Packet.Sender.Name, shipment.Packet.Receiver.Name, shipment.Packet.Destination.Address, strconv.FormatBool(shipment.IsReceived)})
	}

	err = writer.WriteAll(data)
	if err != nil {
		log.Fatal(err)
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}
}
