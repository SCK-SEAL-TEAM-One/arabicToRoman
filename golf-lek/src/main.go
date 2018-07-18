package main

import (
	"api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/convertToRomanNumber", api.ConvertToRomanSymbolHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))

}
