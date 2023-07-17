package handler

import (
	"net/http"
	"weeklytest/domain"
	"weeklytest/helpers"
	"weeklytest/service"

	"github.com/asaskevich/govalidator"
)

func PostPacketHandler(w http.ResponseWriter, r *http.Request) {
	err := helpers.IsPost(w, r)
	if err != nil {
		helpers.FailedResponse(w, 405, "method not allowed")
		return
	}

	var request domain.RequestPacket
	helpers.ReadFromRequestBody(r, &request)

	_, err = govalidator.ValidateStruct(request)
	if err != nil {
		helpers.FailedResponse(w, 400, err.Error())
		return
	}

	packet := service.AddPacket(request)

	helpers.SuccessResponse(w, 201, packet)

}

func GetAllPacketReceivedHandler(w http.ResponseWriter, r *http.Request) {
	err := helpers.IsGet(w, r)
	if err != nil {
		helpers.FailedResponse(w, 405, "method not allowed")
		return
	}

	packet := service.ListPacketReceived()
	helpers.SuccessResponse(w, http.StatusOK, packet)
}

func GetAllPacketsByLocationNameHandler(w http.ResponseWriter, r *http.Request) {
	err := helpers.IsGet(w, r)
	if err != nil {
		helpers.FailedResponse(w, 405, "method not allowed")
		return
	}

	location := r.URL.Query().Get("name")
	packets := service.GetAllPacketsByLocationName(location)

	helpers.SuccessResponse(w, http.StatusOK, packets)
}
