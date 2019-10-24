package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("çıktım ula")
	go func() {
		for {
			fmt.Println("döngü")
		}

	}()
	time.Sleep(2 * time.Second)
}
