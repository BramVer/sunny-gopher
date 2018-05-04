package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", get)
	http.ListenAndServe(":6969", nil)
}

// Defines default GET behaviour
func get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, getQuote())
}
