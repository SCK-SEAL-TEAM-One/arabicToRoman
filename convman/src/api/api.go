package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"romannumber"
	"strconv"
)

type RomanResponse struct {
	RomanNumber string
}

func GetNumber(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	number := query.Get("number")
	convertnumber, _ := strconv.Atoi(number)
	roman := romannumber.ConvertToRoman(convertnumber)
	romanResponse := RomanResponse{RomanNumber: roman}
	romanRosponseString, _ := json.Marshal(romanResponse)
	if number != "" {
		w.Write(romanRosponseString)
		return
	}
	fmt.Fprint(w, "not found number")
}
