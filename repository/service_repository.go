package repository

import (
	"errors"
	"strings"
	"weeklytest/domain"
	"weeklytest/helpers"

	"github.com/google/uuid"
)

type ListService struct {
	Services []domain.Service
}

func NewService(name string, priceKg float64) *domain.Service {
	id := "service-" + uuid.New().String()
	return &domain.Service{Id: id, Name: name, PriceKg: priceKg}
}

func AddService(service domain.Service) {
	listService := ListService{}

	path := "/home/ariwiraa/go/src/phincon/week2/weeklytest/files/service.json"
	helpers.OpenFile(path, &listService)

	listService.Services = append(listService.Services, service)

	helpers.WriteToFileJson(path, listService)
}

func FindServiceByName(name string) (*domain.Service, error) {
	listService := ListService{}

	path := "/home/ariwiraa/go/src/phincon/week2/weeklytest/files/service.json"
	helpers.OpenFile(path, &listService)

	for _, service := range listService.Services {
		if strings.EqualFold(service.Name, name) {
			return &service, nil
		}
	}

	return nil, errors.New("service tidak ditemukan")
}
