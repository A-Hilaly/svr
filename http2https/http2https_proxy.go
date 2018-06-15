package http2https

import (
	"log"
	"net/http"
)

func Default() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		target := "https://" + r.Host + r.URL.Path
		pp := r.URL.Path[1:]
		log.Printf("Pinged. Details:\n  REMOTE ADDR: %s\n  Pattern: %s\n\n", r.RemoteAddr, pp)
		http.Redirect(w, r, target, http.StatusTemporaryRedirect)
	}
}
