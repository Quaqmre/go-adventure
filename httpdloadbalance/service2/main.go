package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "server2")
	})
	http.HandleFunc("/healt", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	log.Println("Server running...")

	http.ListenAndServe(":9000", nil)
}
