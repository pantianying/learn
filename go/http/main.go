package main

import (
	"github.com/didip/tollbooth"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	// Create a request limiter per handler.
	http.Handle("/", tollbooth.LimitFuncHandler(tollbooth.NewLimiter(1, nil), HelloHandler))
	http.ListenAndServe(":12345", nil)
}
