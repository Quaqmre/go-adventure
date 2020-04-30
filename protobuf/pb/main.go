package main

import (
	"encoding/json"
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

type person struct {
	Name  string
	Count int32
}

func main() {
	p := &Person{
		Name:  "tttt",
		Count: 4,
	}

	data, err := proto.Marshal(p)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	fmt.Println("----")

	pp := person{
		Name:  "aaaa",
		Count: 3,
	}

	a, err := json.Marshal(pp)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(a))
}
