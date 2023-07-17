package domain

type Packet struct {
	Id            string   `json:"id"`
	Weight        float64  `json:"weight"`
	StartLocation Location `json:"start_location"`
	Sender        Sender   `json:"sender"`
	Receiver      Receiver `json:"receiver"`
	Destination   Location `json:"destination"`
}

type RequestPacket struct {
	Weight        float64 `json:"weight" validate:"required, alphanumeric"`
	SenderId      string  `json:"sender_id" validate:"required"`
	ReceiverId    string  `json:"receiver_id" validate:"required"`
	StartLocation string  `json:"start_location" validate:"required"`
	Destination   string  `json:"destination" validate:"required"`
}

type PacketDetails struct {
	Packet     Packet `json:"packet"`
	IsReceived bool   `json:"is_received"`
}
