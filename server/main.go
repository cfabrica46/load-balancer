package main

import (
	"fmt"
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

	log.SetFlags(log.Lshortfile)

	http.HandleFunc("/", forwardRequest)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

}

func forwardRequest(w http.ResponseWriter, r *http.Request) {

	URL, err := getServer(r.URL.Path)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rProxy := httputil.NewSingleHostReverseProxy(URL)

	rProxy.ServeHTTP(w, r)
}

func getServer(u string) (URL *url.URL, err error) {

	if lastServerIndex == len(serverList) {
		lastServerIndex = 0
	}

	s := fmt.Sprintf("%s%s", serverList[lastServerIndex], u)

	URL, err = url.Parse(s)

	fmt.Println(URL)
	lastServerIndex++
	return
}
