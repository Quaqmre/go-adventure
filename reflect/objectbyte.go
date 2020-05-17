package main

import "fmt"

type forbyte struct {
	Name string
	Age  int
}

func main() {
	akif := forbyte{"akif", 12}
	acceptinterface(akif)

}
func acceptinterface(in interface{}) {
	byteslice, ok := in.([]byte)
	if ok {
		fmt.Println(byteslice)
	} else {
		fmt.Println("dönüştürülemez")
	}
}
