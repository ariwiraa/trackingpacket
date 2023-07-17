package repository

import (
	"errors"
	"strings"
	"weeklytest/domain"
	"weeklytest/helpers"

	"github.com/google/uuid"
)

type ListPacket struct {
	Packets []domain.Packet `json:"packets"`
}

func NewPacket(weight float64, sender domain.Sender, receiver domain.Receiver, start domain.Location, destination domain.Location) *domain.Packet {
	id := "packet-" + uuid.New().String()
	return &domain.Packet{Id: id, Weight: weight, Sender: sender, Receiver: receiver, StartLocation: start, Destination: destination}
}

func NewPacketDetails(packet domain.Packet, isReceived bool) *domain.PacketDetails {
	return &domain.PacketDetails{Packet: packet, IsReceived: isReceived}
}

func AddPacket(packet domain.Packet) {
	listPacket := ListPacket{}

	path := "/home/ariwiraa/go/src/phincon/week2/weeklytest/files/packet.json"
	helpers.OpenFile(path, &listPacket)

	listPacket.Packets = append(listPacket.Packets, packet)

	helpers.WriteToFileJson(path, listPacket)
}

func FindPacketById(id string) (*domain.Packet, error) {
	listPacket := ListPacket{}

	path := "/home/ariwiraa/go/src/phincon/week2/weeklytest/files/packet.json"
	helpers.OpenFile(path, &listPacket)

	for _, packet := range listPacket.Packets {
		if strings.EqualFold(packet.Id, id) {
			return &packet, nil
		}
	}

	return nil, errors.New("id paket tidak ditemukan")
}
