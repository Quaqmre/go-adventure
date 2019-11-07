package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	data := os.Args[1:]
	data1 := []int{}
	for _, v := range data {
		i, _ := strconv.Atoi(v)
		data1 = append(data1, i)
	}
	a1, b1, c1, d1, e1, f1 := fun_degree(data1)
	fmt.Println(a1, b1, c1, d1, e1, f1)

}

func fun_degree(data []int) (a, b, c, d, e, f []int) {
	for i := 0; i < len(data); i++ {
		value := data[i]
		if value >= 90 && value <= 100 {
			a = append(a, value)
		}
		if value >= 80 && value < 90 {
			b = append(b, value)
		}
		if value >= 70 && value < 80 {
			c = append(c, value)
		}
		if value >= 60 && value < 70 {
			d = append(d, value)
		}
		if value >= 50 && value < 60 {
			e = append(e, value)
		}
		if value < 50 {
			f = append(f, value)
		}
	}
	return a, b, c, d, e, f
}
