package repository

import (
	"strings"
	"weeklytest/domain"
	"weeklytest/helpers"

	"github.com/google/uuid"
)

type ListLocation struct {
	Locations []domain.Location `json:"locations"`
}

func NewLocation(name, address string) *domain.Location {
	id := "loc-" + uuid.New().String()
	return &domain.Location{Id: id, Name: name, Address: address}
}

func AddLocation(location domain.Location) {
	listLocation := ListLocation{}

	path := "/home/ariwiraa/go/src/phincon/week2/weeklytest/files/location.json"
	helpers.OpenFile(path, &listLocation)

	listLocation.Locations = append(listLocation.Locations, location)

	helpers.WriteToFileJson(path, listLocation)
}

func GetAllLocations() []domain.Location {
	listLocation := ListLocation{}

	path := "/home/ariwiraa/go/src/phincon/week2/weeklytest/files/location.json"
	helpers.OpenFile(path, &listLocation)

	return listLocation.Locations
}

func FindLocationByName(name string) *domain.Location {
	listLocation := ListLocation{}

	path := "/home/ariwiraa/go/src/phincon/week2/weeklytest/files/location.json"
	helpers.OpenFile(path, &listLocation)

	for _, loc := range listLocation.Locations {
		if strings.EqualFold(loc.Name, name) {
			return &loc
		}
	}

	return nil
}

func FindAllLocations() []domain.Location {
	listLocation := ListLocation{}

	path := "/home/ariwiraa/go/src/phincon/week2/weeklytest/files/location.json"
	helpers.OpenFile(path, &listLocation)

	return listLocation.Locations
}
