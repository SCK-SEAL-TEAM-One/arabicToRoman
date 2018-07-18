package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"romannumber"
)

type RomanResponse struct {
	RomanNumber string `json:"romanNumber"`
}

func NumberToRomanNumberHandler(writer http.ResponseWriter, request *http.Request) {
	numberString := request.URL.Query().Get("number")
	number, _ := strconv.Atoi(numberString)

	roman := romannumber.ConvertToRomanNumber(number)
	romanJson, _ := json.Marshal(RomanResponse{
		RomanNumber: roman,
	})
	writer.Write(romanJson)
}
