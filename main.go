package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	version = "1"
	addr    string
)

func init() {
	flag.StringVar(&addr, "addr", ":8080", "Address to listen on")
	flag.StringVar(&version, "version", version, "Version of the server")
}

func main() {
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		logReq(r)
		w.Header().Add("Content-Type", "text/plain")
		fmt.Fprintln(w, version)
	})
	mux.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
		logReq(r)
		w.Header().Add("Content-Type", "text/plain")
		for _, env := range os.Environ() {
			if _, err := fmt.Fprintln(w, env); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	})
	log.Fatal(http.ListenAndServe(addr, mux))
}

func logReq(r *http.Request) {
	log.Printf(`method=%s path=%s fwd=%s host=%s`, r.Method, r.RequestURI, r.RemoteAddr, r.Host)
}
