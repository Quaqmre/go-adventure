package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var i uint32 = 1
	b := atomic.CompareAndSwapUint32(&i, 2, 3)
	fmt.Println(b, i)
	defer func() {
		fmt.Println("1i")
	}()
	defer func() {
		fmt.Println("2i")
	}()
}
