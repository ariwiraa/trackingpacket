package service

import (
	"weeklytest/domain"
	"weeklytest/repository"
)

func AddService(name string, priceKg float64) *domain.Service {
	service := repository.NewService(name, priceKg)

	repository.AddService(*service)

	return service
}
