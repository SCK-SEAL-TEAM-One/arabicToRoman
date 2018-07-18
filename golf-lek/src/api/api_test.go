package api_test

import (
	. "api"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_API_ConvertToRomanSymbolHandler_Input_1_Should_Be_I(t *testing.T) {
	expectedResponse := RomanResponse{
		RomanNumber: "I",
	}
	req := httptest.NewRequest("GET", "localhost:3000/convertToRomanNumber?number=1", nil)
	w := httptest.NewRecorder()
	ConvertToRomanSymbolHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var actual RomanResponse
	json.Unmarshal(body, &actual)

	if expectedResponse != actual {
		t.Errorf("expected is %v but got %v", expectedResponse, actual)
	}
}
