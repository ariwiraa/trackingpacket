package handler

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"weeklytest/domain"
	"weeklytest/helpers"
	"weeklytest/service"

	"github.com/asaskevich/govalidator"
)

func PostShipmentHandler(w http.ResponseWriter, r *http.Request) {
	err := helpers.IsPost(w, r)
	if err != nil {
		helpers.FailedResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var request domain.RequestShipment
	helpers.ReadFromRequestBody(r, &request)

	_, err = govalidator.ValidateStruct(request)
	if err != nil {
		helpers.FailedResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	shipment, err := service.AddShipment(request)
	if err != nil {
		helpers.FailedResponse(w, http.StatusBadRequest, err.Error())
	}

	helpers.SuccessResponse(w, http.StatusCreated, shipment)
}

func PostUpdateCheckPointHandler(w http.ResponseWriter, r *http.Request) {
	err := helpers.IsPut(w, r)
	if err != nil {
		helpers.FailedResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var request domain.RequestUpdateShipment
	helpers.ReadFromRequestBody(r, &request)

	_, err = govalidator.ValidateStruct(request)
	if err != nil {
		helpers.FailedResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	shipment := service.UpdateCheckPoint(request.ShipmentId, request.LocationName)

	helpers.SuccessResponse(w, http.StatusCreated, shipment)
}

func GetShipmentByIdHandler(w http.ResponseWriter, r *http.Request) {
	err := helpers.IsGet(w, r)
	if err != nil {
		helpers.FailedResponse(w, 405, "method not allowed")
		return
	}

	id := r.URL.Query().Get("id")
	shipment, err := service.GetShipmentById(id)
	if err != nil {
		helpers.FailedResponse(w, http.StatusNotFound, err.Error())
	}

	helpers.SuccessResponse(w, http.StatusOK, shipment)
}

func DownloadFileCSV(w http.ResponseWriter, r *http.Request) {
	service.CreateFileCSV()

	path := "/home/ariwiraa/go/src/phincon/week2/weeklytest/files/shipment.csv"
	file, err := os.Open(path)
	if err != nil {
		helpers.FailedResponse(w, http.StatusNotFound, "file not found")
	}
	defer file.Close()

	fileInfo, _ := file.Stat()

	w.Header().Set("Content-Disposition", "attachment; filename="+fileInfo.Name())
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

	io.Copy(w, file)
}
