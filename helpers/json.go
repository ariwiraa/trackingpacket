package helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()
	err := decoder.Decode(result)
	if err != nil {
		panic(err)
	}
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	if err != nil {
		panic(err)
	}
}

func WriteToFileJson(path string, data interface{}) {
	writer, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	encoder := json.NewEncoder(writer)
	encoder.Encode(data)
}

func OpenFile(path string, data interface{}) {

	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	if len(file) == 0 {
		return
	}

	err = json.Unmarshal(file, data)
	if err != nil {
		panic(err)
	}
}
