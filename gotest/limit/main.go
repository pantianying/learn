package main

import (
	"github.com/throttled/throttled"
	"github.com/throttled/throttled/store/memstore"
	"log"
	"net/http"
)

func main() {
	store, err := memstore.New(65536)
	if err != nil {
		log.Fatal(err)
	}

	quota := throttled.RateQuota{throttled.PerMin(20), 5}
	rateLimiter, err := throttled.NewGCRARateLimiter(store, quota)
	if err != nil {
		log.Fatal(err)
	}

	httpRateLimiter := throttled.HTTPRateLimiter{
		RateLimiter: rateLimiter,
		VaryBy:      &throttled.VaryBy{Path: true},
	}

	mux := http.NewServeMux()
	mux.Handle("/", httpRateLimiter.RateLimit(&myHandler{}))
	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	srv.ListenAndServe()

}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is version 3"))
}
