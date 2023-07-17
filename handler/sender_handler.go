package handler

import (
	"net/http"
	"weeklytest/domain"
	"weeklytest/helpers"
	"weeklytest/service"

	"github.com/asaskevich/govalidator"
)

func PostSenderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

	var request domain.RequestSender
	helpers.ReadFromRequestBody(r, &request)

	_, err := govalidator.ValidateStruct(request)
	if err != nil {
		helpers.FailedResponse(w, 400, err.Error())
		return
	}

	sender := service.AddSender(request.Name, request.Phone)

	helpers.SuccessResponse(w, 201, sender)
}

func GetByIdSenderHandler(w http.ResponseWriter, r *http.Request) {
	err := helpers.IsGet(w, r)
	if err != nil {
		helpers.FailedResponse(w, 405, "method not allowed")
	}

	idSender := r.URL.Query().Get("id")

	sender := service.FindSenderById(idSender)

	helpers.SuccessResponse(w, 200, sender)
}
