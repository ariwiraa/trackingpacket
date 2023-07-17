package domain

type Sender struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type RequestSender struct {
	Name  string `json:"name" valid:"required"`
	Phone string `json:"phone" valid:"required, numeric"`
}
