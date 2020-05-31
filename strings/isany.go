package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "qwes"
	bigs := "aaaaaaaaaaaaaaaaaas"
	b := strings.IndexAny(bigs, s)
	fmt.Println(b)
}
