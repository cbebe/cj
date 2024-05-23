package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func deepCompare(sf, df io.Reader) bool {
	sscan := bufio.NewScanner(sf)
	dscan := bufio.NewScanner(df)

	for sscan.Scan() {
		dscan.Scan()
		if !bytes.Equal(sscan.Bytes(), dscan.Bytes()) {
			return true
		}
	}

	return false
}

var local url.URL
var server http.Handler

func proxy(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.Join(strings.Split(r.URL.Path, "/")[2:], "/")
		r.Host = local.Host
		p.ServeHTTP(w, r)
	}
}

func renderIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/index.html" || r.URL.Path == "/" {
		app().Render(context.Background(), w)
	} else {
		server.ServeHTTP(w, r)
	}
}

func serve() error {
	if err := os.Chdir("docs"); err != nil {
		return err
	}
	local, err := url.Parse("http://localhost:8080")
	if err != nil {
		return err
	}

	rp := httputil.NewSingleHostReverseProxy(local)
	http.HandleFunc("/cj/", proxy(rp))
	server = http.FileServer(http.Dir("."))
	http.Handle("/", http.HandlerFunc(renderIndex))
	return http.ListenAndServe("0.0.0.0:8080", nil)
}

func updateHtml() error {
	b, err := os.ReadFile("docs/index.html")
	if err != nil {
		if _, ok := err.(*os.PathError); !ok {
			return err
		} else {
			b = []byte{}
		}
	}
	appHTML := &bytes.Buffer{}
	err = app().Render(context.Background(), appHTML)
	if err != nil {
		return err
	}
	bufore := bytes.NewBuffer(b)
	bufter := bytes.NewBuffer(bytes.Clone(appHTML.Bytes()))
	if deepCompare(bufore, bufter) {
		err := os.WriteFile("docs/index.html", appHTML.Bytes(), 0o644)
		if err != nil {
			return fmt.Errorf("error writing to file: %v", err)
		}
	}
	return nil
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "build" {
		if err := updateHtml(); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := serve(); err != nil {
			log.Fatal(err)
		}
	}
}
