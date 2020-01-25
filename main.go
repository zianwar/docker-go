package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

var addr *string = flag.String("addr", "127.0.0.1:3001", "http server address")

func main() {
	flag.Parse()
	log.SetHandler(text.New(os.Stderr))

	http.HandleFunc("/", loggingMiddleware(handle))

	log.Infof("http server listening on http://%s\n", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatalf("%v", err)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func loggingMiddleware(h http.HandlerFunc) http.HandlerFunc {
	count := 0
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		count++
		log.WithFields(log.Fields{
			"URI":        r.RequestURI,
			"User-Agent": r.UserAgent(),
			"IP":         r.RemoteAddr,
		}).Infof("Request #%d\n", count)
		h.ServeHTTP(w, r)
	})
}
