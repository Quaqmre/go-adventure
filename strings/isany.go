package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "qwes"
	bigs := "aaaaaaaaaaaaaaaaaas" //input lengt 8 den büyükse 128 karakter için 8 bytlık 8 lenghtlik bir array aoluşturuğ 128 bit
	//set edilerek olan harfleri bu bitlere set ediliyor
	b := strings.IndexAny(bigs, s)
	fmt.Println(b)
}
