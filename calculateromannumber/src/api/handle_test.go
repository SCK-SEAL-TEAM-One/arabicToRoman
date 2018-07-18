package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_API_NumberToRomanNumberHandler_Input_1_Should_be_convertToRomanNumberResponse(t *testing.T) {
	url := "/convertToRomanNumber?number=1"
	request := httptest.NewRequest("GET", url, nil)
	responseRecorder := httptest.NewRecorder()
	expectedResponse := RomanResponse{
		RomanNumber: "I",
	}
	NumberToRomanNumberHandler(responseRecorder, request)
	response := responseRecorder.Result()
	body, _ := ioutil.ReadAll(response.Body)

	var actual RomanResponse
	json.Unmarshal(body, &actual)

	if expectedResponse != actual {
		t.Errorf("Should Be %v but it got %v", expectedResponse, actual)
	}
}
