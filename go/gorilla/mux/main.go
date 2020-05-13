package main

import (
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Path("/v1.0/apt/devices/{device_id}/locations")
	r.Methods("POST")
	r.Schemes("http")

	//debug.Debug(r.Match(http.Request{}))
}
