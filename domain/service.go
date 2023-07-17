package domain

type Service struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	PriceKg float64 `json:"price_kg"`
}

type RequestService struct {
	Name    string  `json:"name" validate:"required"`
	PriceKg float64 `json:"price_kg" validate:"required, numeric"`
}
