package services

import "fmt"

var senderDatas = []Sender{}

type Sender struct {
	Id    string
	Name  string
	phone string
}

func AddSender(name, phone string) *Sender {
	s := NewSender()
	s.Name = name
	s.phone = phone

	senderDatas = append(senderDatas, *s)
	fmt.Println("pengirim sudah terdaftar")

	return s
}

func NewSender() *Sender {
	return &Sender{}
}
