package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := [3]int{}
	_ = copy(b, a)
	b[2] = 5
	fmt.Println(b)
	fmt.Println(a)
}
