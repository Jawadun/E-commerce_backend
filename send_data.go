package main

import (
	"encoding/json"
	"net/http"
)

func sendData(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
