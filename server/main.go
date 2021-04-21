package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	serverList = []string{
		"http://localhost:8001",
		"http://localhost:8002",
		"http://localhost:8003",
		"http://localhost:8004",
	}
	lastServerIndex = 0
)

func main() {

	http.HandleFunc("/", forwardRequest)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

}

func forwardRequest(w http.ResponseWriter, r *http.Request) {

	URL, err := getServer()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rProxy := httputil.NewSingleHostReverseProxy(URL)

	rProxy.ServeHTTP(w, r)
}

func getServer() (URL *url.URL, err error) {

	if lastServerIndex == len(serverList) {
		lastServerIndex = 0
	}

	URL, err = url.Parse(serverList[lastServerIndex])
	lastServerIndex++
	return
}
