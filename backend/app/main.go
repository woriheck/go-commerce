package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/phone", HelloPhone)
	http.ListenAndServe(":8080", nil)
}

type Response struct {
	Message string `json:"message"`
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	jsonOut, _ := json.Marshal(Response{Message: "Hello Product"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", jsonOut)
}

func HelloPhone(w http.ResponseWriter, r *http.Request) {
	jsonOut, _ := json.Marshal(Response{Message: "Hello Phone"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", jsonOut)
}
