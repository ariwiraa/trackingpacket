package handler

import (
	"net/http"
	"weeklytest/domain"
	"weeklytest/helpers"
	"weeklytest/service"

	"github.com/asaskevich/govalidator"
)

func PostLocationHandler(w http.ResponseWriter, r *http.Request) {
	err := helpers.IsPost(w, r)
	if err != nil {
		helpers.FailedResponse(w, 405, "method not allowed")
		return
	}

	var request domain.RequestLocation
	helpers.ReadFromRequestBody(r, &request)

	_, err = govalidator.ValidateStruct(request)
	if err != nil {
		helpers.FailedResponse(w, 400, err.Error())
		return
	}

	location := service.AddLocation(request.Name, request.Address)

	helpers.SuccessResponse(w, 201, location)

}

func GetLocationHandler(w http.ResponseWriter, r *http.Request) {
	err := helpers.IsGet(w, r)
	if err != nil {
		helpers.FailedResponse(w, 405, "method not allowed")
		return
	}

	locations := service.GetAllLocations()

	helpers.SuccessResponse(w, 200, locations)
}
