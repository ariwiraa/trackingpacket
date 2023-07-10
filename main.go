package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ariwiraa/trackingpacket/services"
	"github.com/manifoldco/promptui"
)

func main() {

	addDefaultService()
	addDefaultLocation()
	serviceNames := services.GetAllServicesName()

Menu:
	for {
		prompt := promptui.Select{Label: "Menu", Items: []string{"Exit", "Shipment", "List checkpoint", "list all packet received", "List packet by checkpoint", "update shipment checkpoint"}}
		cmd, _, _ := prompt.Run()

		switch cmd {
		case 1:
			createShipment(serviceNames)
		case 2:
			GetAllLocationsByName()
		case 3:
			listPacketReceived()
		case 4:
			getAllPacketsByLocationName()
		case 5:
			updateCheckPoint()
		case 0:
			break Menu
		}
	}
}

func listPacketReceived() {
	packet := services.ListPacketReceived()
	for _, v := range packet {
		fmt.Println("========= PAKET DITERIMA ============")
		fmt.Printf("Pengirim: %v\n", v.Packet.Sender.Name)
		fmt.Printf("Penerima: %v\n", v.Packet.Sender.Name)
		fmt.Printf("status: %v\n", isReceived(v.IsReceived))
		fmt.Println("=====================")
	}
}

func createShipment(servicesName []string) {
	prompt := promptui.Prompt{Label: "Nama Pengirim", Validate: notNull}
	senderName, err := prompt.Run()
	panicIfError(err)

	prompt = promptui.Prompt{Label: "No Telepon Pengirim", Validate: notNull}
	senderPhone, err := prompt.Run()
	panicIfError(err)
	inputSender := services.AddSender(senderName, senderPhone)

	prompt = promptui.Prompt{Label: "Nama Penerima", Validate: notNull}
	receiverName, err := prompt.Run()
	panicIfError(err)

	prompt = promptui.Prompt{Label: "No Telepon Penerima", Validate: notNull}
	receiverPhone, err := prompt.Run()
	panicIfError(err)
	inputReceiver := services.AddReceiver(receiverName, receiverPhone)

	prompt = promptui.Prompt{Label: "Nama Lokasi Tujuan", Validate: notNull}
	locationName, err := prompt.Run()
	panicIfError(err)

	prompt = promptui.Prompt{Label: "Alamat Tujuan", Validate: notNull}
	locationAddress, err := prompt.Run()
	panicIfError(err)
	inputLocation := services.AddLocation(locationName, locationAddress)

	promptSelect := promptui.Select{Label: "Nama services", Items: servicesName}
	_, serviceName, err := promptSelect.Run()
	panicIfError(err)
	inputService := services.AddServiceToPacket(serviceName)

	prompt = promptui.Prompt{Label: "Berat Paket", Validate: notNull}
	weight, err := prompt.Run()
	panicIfError(err)
	convertWeight, err := strconv.Atoi(weight)
	panicIfError(err)

	packet := services.AddPacket(inputSender, inputReceiver, *inputLocation, float64(convertWeight))
	services.AddDelivery(packet, inputService)
	fmt.Println("CREATED")
}

func GetAllLocationsByName() {

	locations := services.GetAllLocations()
	for _, loc := range locations {
		if strings.Contains(loc.Name, "Gudang") {
			fmt.Printf("Nama Gudang: %s\n", loc.Name)
		}

	}

}

func getAllPacketsByLocationName() {
	prompt := promptui.Prompt{Label: "Nama lokasi", Validate: notNull}
	locations, err := prompt.Run()
	panicIfError(err)

	packet := services.GetAllPacketsByLocationName(locations)
	for _, v := range packet {
		fmt.Printf("ID Packet: %s", v.Id)
	}
	fmt.Println(packet)

}

func updateCheckPoint() {
	prompt := promptui.Prompt{Label: "Id Pengiriman"}
	deliveryId, err := prompt.Run()
	panicIfError(err)

	prompt = promptui.Prompt{Label: "Id Lokasi"}
	locationId, err := prompt.Run()
	panicIfError(err)

	services.UpdateCheckPoint(deliveryId, locationId)
}

func addDefaultService() {
	regularService := services.Service{
		Id:      "1",
		Name:    "Regular",
		PriceKg: 11000,
	}

	CargoService := services.Service{
		Id:      "2",
		Name:    "Cargo",
		PriceKg: 22000,
	}

	services.AddService(regularService)
	services.AddService(CargoService)
}

func addDefaultLocation() {
	services.AddLocation("Gudang asik", "gudang no 7")
	services.AddLocation("Gudang uhuy", "gudang no 8")
	services.AddLocation("Gudang gnadug", "gudang no 9")
}

func isReceived(input bool) string {
	result := "belum diterima"
	if input {
		result = "diterima"
	}

	return result
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func notNull(input string) error {
	if len(input) == 0 {
		return errors.New("tidak boleh kosong")
	}

	return nil
}
