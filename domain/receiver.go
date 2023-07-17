package domain

type Receiver struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type RequestReceiver struct {
	Name  string `json:"name" valid:"required"`
	Phone string `json:"phone" valid:"required,numeric"`
}
