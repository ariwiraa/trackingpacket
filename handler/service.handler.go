package handler

import (
	"net/http"
	"weeklytest/domain"
	"weeklytest/helpers"
	"weeklytest/service"

	"github.com/asaskevich/govalidator"
)

func PostServiceHandler(w http.ResponseWriter, r *http.Request) {
	err := helpers.IsPost(w, r)
	if err != nil {
		helpers.FailedResponse(w, 405, "method not allowed")
		return
	}

	var request domain.RequestService
	helpers.ReadFromRequestBody(r, &request)

	_, err = govalidator.ValidateStruct(request)
	if err != nil {
		helpers.FailedResponse(w, 400, err.Error())
		return
	}

	service := service.AddService(request.Name, request.PriceKg)

	helpers.SuccessResponse(w, 201, service)
}
