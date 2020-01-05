package main

import (
	"container/ring"
	"fmt"
)

type as struct{}

func main() {
	a := []string{"akif", "ali", "ayşe"}
	// ring oluşturuyor ve sonsuza kadar dönmektedir,
	// ring mantığı eğer tek instance warsa next during vbe prev hep kendidir.

	ring := ring.New(len(a))
	for i := 0; i < len(a); i++ {
		ring.Value = a[i]
		ring = ring.Next()
	}

	for {
		ring = ring.Next()
		fmt.Println(ring.Value)
	}

	// b := new(as)
	// t := b
	// fmt.Println(t)

}
