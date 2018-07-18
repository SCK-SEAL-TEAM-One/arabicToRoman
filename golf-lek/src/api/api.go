package api

import (
	"encoding/json"
	"net/http"
	"romannumber"
	"strconv"
)

type RomanResponse struct {
	RomanNumber string `json:"romanNumber"`
}

func ConvertToRomanSymbolHandler(writer http.ResponseWriter, request *http.Request) {
	numberString := request.URL.Query().Get("number")
	number, _ := strconv.Atoi(numberString)

	roman := romannumber.ConvertToRomanSymbol(number)
	romanJson, _ := json.Marshal(RomanResponse{
		RomanNumber: roman,
	})
	writer.Write(romanJson)
}
