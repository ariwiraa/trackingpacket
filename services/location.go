package services

import (
	"fmt"
	"strconv"
	"strings"
)

var locations []Location

// Lokasi
type Location struct {
	Id     string
	Name   string
	Addres string
}

func NewLocation() *Location {
	return &Location{}
}

func AddLocation(name, address string) *Location {
	l := NewLocation()
	id := len(locations) + 1

	l.Id = strconv.Itoa(id)
	l.Name = name
	l.Addres = address
	locations = append(locations, *l)

	return l
}

func GetAllLocations() []Location {
	return locations
}

func FindLocationById(id string) *Location {
	found := false
	for i, loc := range locations {
		if strings.EqualFold(loc.Id, id) {
			found = true
			return &locations[i]
		}
	}

	if !found {
		fmt.Println("data tidak ditemukan")
	}

	return nil
}
