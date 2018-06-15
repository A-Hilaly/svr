package httpproxy

import (
	"log"
	"net/http"
)

func Handler(url string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pp := r.URL.Path[1:]
		log.Printf("Pinged. Details:\n  REMOTE ADDR: %s\n  Pattern: %s\n\n", r.RemoteAddr, pp)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
	}
}

func Default() func(http.ResponseWriter, *http.Request) {
	return Handler("https://google.com/")
}
