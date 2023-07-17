package service

import (
	"weeklytest/domain"
	"weeklytest/repository"
)

func AddLocation(name, address string) *domain.Location {
	location := repository.NewLocation(name, address)

	repository.AddLocation(*location)

	return location
}

func GetAllLocations() []domain.Location {
	return repository.GetAllLocations()
}
