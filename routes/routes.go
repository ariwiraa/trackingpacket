package routes

import (
	"net/http"
	"weeklytest/handler"
)

func Routes(mux *http.ServeMux) {
	// r := &http.Request{}
	senderRoute(mux)
	receiverRoutes(mux)
	locationRoutes(mux)
	serviceRoutes(mux)
	packetRoutes(mux)
	shipmentRoutes(mux)
}

func shipmentRoutes(mux *http.ServeMux) {

	var postShipmentHandler http.HandlerFunc = handler.PostShipmentHandler
	var updateCheckPointHandler http.HandlerFunc = handler.PostUpdateCheckPointHandler
	var getShipmentById http.HandlerFunc = handler.GetShipmentByIdHandler
	var download http.HandlerFunc = handler.DownloadFileCSV

	mux.Handle("/shipment", postShipmentHandler)
	mux.Handle("/shipment/update", updateCheckPointHandler)
	mux.Handle("/shipment/search", getShipmentById)
	mux.Handle("/shipment/download", download)

}

func packetRoutes(mux *http.ServeMux) {
	var postPacketHandler http.HandlerFunc = handler.PostPacketHandler
	var getAllPacketReceivedHandler http.HandlerFunc = handler.GetAllPacketReceivedHandler
	var getAllPacketsByLocationNameHandler http.HandlerFunc = handler.GetAllPacketsByLocationNameHandler

	mux.Handle("/packet", postPacketHandler)
	mux.Handle("/packet/received", getAllPacketReceivedHandler)
	mux.Handle("/packet/search", getAllPacketsByLocationNameHandler)
}

func serviceRoutes(mux *http.ServeMux) {
	var postServiceHandler http.HandlerFunc = handler.PostServiceHandler

	mux.Handle("/service", postServiceHandler)
}

func locationRoutes(mux *http.ServeMux) {
	var postLocationHandler http.HandlerFunc = handler.PostLocationHandler
	var getLocationHandler http.HandlerFunc = handler.GetLocationHandler

	mux.Handle("/location", postLocationHandler)
	mux.Handle("/locations", getLocationHandler)
}

func receiverRoutes(mux *http.ServeMux) {
	var postReceiverHandler http.HandlerFunc = handler.PostReceiverHandler

	mux.Handle("/receiver", postReceiverHandler)
}

func senderRoute(mux *http.ServeMux) {
	var postSenderHandler http.HandlerFunc = handler.PostSenderHandler
	var getByIdSenderHandler http.HandlerFunc = handler.GetByIdSenderHandler

	mux.Handle("/sender", postSenderHandler)
	mux.Handle("/sender/search", getByIdSenderHandler)
}
