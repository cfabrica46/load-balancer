package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hola", hola1)
	http.HandleFunc("/adios", adios)

	err := http.ListenAndServe(":8003", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func hola1(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hola-3\n")

}

func adios(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Adios-3\n")

}
