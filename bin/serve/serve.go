package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	local, err := url.Parse("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL)
			r.URL.Path = strings.Join(strings.Split(r.URL.Path, "/")[2:], "/")
			r.Host = local.Host
			p.ServeHTTP(w, r)
		}
	}
	proxy := httputil.NewSingleHostReverseProxy(local)
	http.HandleFunc("/cj/", handler(proxy))
	http.Handle("/", http.FileServer(http.Dir(".")))
	err = http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
