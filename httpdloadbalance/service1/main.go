package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "server1")
	})
	http.HandleFunc("/healt", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	log.Println("Server running...")

	http.ListenAndServe(":9001", nil)
}
