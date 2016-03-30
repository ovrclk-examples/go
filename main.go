package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
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
	//panic("fail")
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		logReq(r)
		fmt.Fprintln(w, version)
	})
	log.Fatal(http.ListenAndServe(addr, mux))
}

func logReq(r *http.Request) {
	log.Printf(`method=%s " path=%s fwd="%shost=%s`, r.Method, r.RequestURI, r.RemoteAddr, r.Host)
}
