package domain

type Location struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type RequestLocation struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
}
