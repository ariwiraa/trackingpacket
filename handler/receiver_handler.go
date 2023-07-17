package handler

import (
	"net/http"
	"weeklytest/domain"
	"weeklytest/helpers"
	"weeklytest/service"

	"github.com/asaskevich/govalidator"
)

func PostReceiverHandler(w http.ResponseWriter, r *http.Request) {
	// Cek apakah ini method post atau bukan
	helpers.IsPost(w, r)

	var request domain.RequestReceiver
	helpers.ReadFromRequestBody(r, &request)

	_, err := govalidator.ValidateStruct(request)
	if err != nil {
		helpers.FailedResponse(w, 400, err.Error())
		return
	}

	receiver := service.AddReceiver(request.Name, request.Phone)

	helpers.SuccessResponse(w, 201, receiver)

}
