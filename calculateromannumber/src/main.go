package main

import (
	"log"
	"net/http"

	"./api"
)

func main() {
	http.HandleFunc("/convertToRomanNumber", api.NumberToRomanNumberHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
