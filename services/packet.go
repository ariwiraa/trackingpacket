package services

import (
	"strconv"
)

var packetDatas = []Packet{}

type Packet struct {
	Id          string
	Sender      Sender
	Receiver    Receiver
	Weight      float64
	Destination Location
}

type PacketDetails struct {
	Packet     Packet
	IsReceived bool
}

func AddPacket(sender *Sender, receiver *Receiver, location Location, weight float64) *Packet {
	id := len(packetDatas) + 1
	idPacket := strconv.Itoa(id)

	p := NewPacket(idPacket, *sender, *receiver, weight, location)
	packetDatas = append(packetDatas, *p)

	return p
}

func NewPacket(id string, sender Sender, receiver Receiver, weight float64, destination Location) *Packet {
	return &Packet{Id: id, Sender: sender, Receiver: receiver, Weight: weight, Destination: destination}
}

func NewPacketDetails(packet Packet, isReceived bool) *PacketDetails {
	return &PacketDetails{Packet: packet, IsReceived: isReceived}
}
