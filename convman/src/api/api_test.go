package api_test

import (
	"api"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_GetNumber_Method_Get_Input_1_Should_Be_I(t *testing.T) {
	req := httptest.NewRequest("GET", "/ConvertToRomanNumber?number=1", nil)
	w := httptest.NewRecorder()
	api.GetNumber(w, req)
	expected := api.RomanResponse{RomanNumber: "I"}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	actual := api.RomanResponse{}
	json.Unmarshal(body, &actual)

	if expected != actual {
		t.Errorf("should be %v but is %v", expected, actual)
	}
}
