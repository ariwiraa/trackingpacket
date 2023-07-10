package services

import "strconv"

var receivers = []Receiver{}

type Receiver struct {
	Id    string
	Name  string
	Phone string
}

func AddReceiver(name, phone string) *Receiver {
	id := len(receivers) + 1

	idReceiver := strconv.Itoa(id)

	r := NewReceiver(idReceiver, name, phone)
	receivers = append(receivers, *r)

	return r
}

func NewReceiver(id, name, phone string) *Receiver {
	return &Receiver{Id: id, Name: name, Phone: phone}
}
