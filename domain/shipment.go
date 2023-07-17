package domain

type Shipment struct {
	Id         string     `json:"id"`
	Price      float64    `json:"price"`
	IsReceived bool       `json:"is_received"`
	Service    Service    `json:"service"`
	Packet     Packet     `json:"packet"`
	CheckPoint []Location `json:"checkpoint"`
}

type RequestShipment struct {
	ServiceName string `json:"service_name" validate:"required"`
	PacketId    string `json:"packet_id" validate:"required"`
}

type RequestUpdateShipment struct {
	ShipmentId   string `json:"shipment_id" validate:"required"`
	LocationName string `json:"location_name" validate:"required"`
}
