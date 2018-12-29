package httpserver

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	pp := r.URL.Path[1:]
	log.Printf("Pinged. Details:\n  REMOTE ADDR: %s\n  Pattern: %s\n\n", r.RemoteAddr, pp)
	fmt.Fprintf(w, "SVG V1: Simple http server. Detected pattern: \"%s\".", pp)
}

// New start a new http server
func New(port int, pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	mux := http.NewServeMux()
	mux.HandleFunc(pattern, handler)
	portstr := ":" + strconv.Itoa(port)

	log.Printf("Listening on port %d (pattern: \"%s\")\n", port, pattern)
	err := http.ListenAndServe(portstr, mux)
	if err != nil {
		log.Panic(err)
	}
}

// Default start a default http server
func Default(port int, pattern string) {
	New(port, pattern, defaultHandler)
}
