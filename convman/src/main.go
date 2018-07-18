package main

import (
	"api"
	"net/http"
)

func main() {
	http.HandleFunc("/ConvertToRomanNumber", api.GetNumber)
	http.ListenAndServe(":3000", nil)
}
