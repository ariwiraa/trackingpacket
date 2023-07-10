package services

import (
	"fmt"
	"strings"
)

var services = []Service{}

type Service struct {
	Id      string
	Name    string
	PriceKg float64
}

// func NewService(id, name string, priceKg float64) *Service {
// 	return &Service{Id: id, Name: name, PriceKg: priceKg}
// }

func AddService(service Service) {
	services = append(services, service)
}

func AddServiceToPacket(name string) *Service {
	found := false
	for i, service := range services {
		if strings.EqualFold(service.Name, name) {
			found = true
			return &services[i]
		}
	}

	if !found {
		fmt.Println("data tidak ditemukan")
	}

	return nil
}

func GetAllServicesName() []string {
	var serviceNames []string
	for _, v := range services {
		serviceNames = append(serviceNames, v.Name)
	}

	return serviceNames
}

func DefaultService() {
	regularService := Service{
		Id:      "1",
		Name:    "Regular",
		PriceKg: 11000,
	}

	CargoService := Service{
		Id:      "2",
		Name:    "Cargo",
		PriceKg: 22000,
	}

	AddService(regularService)
	AddService(CargoService)
}
