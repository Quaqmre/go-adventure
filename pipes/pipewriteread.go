package main

import (
	"fmt"
	"io"
)

func main() {
	r, w := io.Pipe()
	go func() {
		for {

			fmt.Fprint(w, "test")
		}
	}()
	for {
		a := make([]byte, 500)
		size, _ := r.Read(a)
		str := a[:size]
		fmt.Println(string(str))
	}
}
