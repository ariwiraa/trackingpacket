package helpers

import "net/http"

type ResponseWithData struct {
	Code   int
	Status string
	Data   interface{}
}

type ResponseWithMessage struct {
	Code    int
	Status  string
	Message string
}

func SuccessResponse(w http.ResponseWriter, code int, data interface{}) {
	response := ResponseWithData{
		Code:   code,
		Status: "Success",
		Data:   data,
	}

	WriteToResponseBody(w, response)
}

func FailedResponse(w http.ResponseWriter, code int, message string) {
	response := ResponseWithMessage{
		Code:    code,
		Status:  "failed",
		Message: message,
	}

	WriteToResponseBody(w, response)
}
