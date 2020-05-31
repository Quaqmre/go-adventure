package main

import "fmt"

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}
func main() {
	s := make([]byte, 30)
	slice := s[10:20]

	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}
	fmt.Println("before", slice)
	AddOneToEachElement(slice)
	fmt.Println("after", slice)
	fmt.Println("s", s)
}
