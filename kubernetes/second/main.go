package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	a, _ := http.Get("http://172.17.0.2:4000")
	defer a.Body.Close()
	as, _ := ioutil.ReadAll(a.Body)

	f, _ := os.OpenFile("asd")

	str := string(as)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ikinci: %s", str)
	})
	http.ListenAndServe(":5000", nil)

}
