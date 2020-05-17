package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// scanner := bufio.NewScanner(file)
	r := bufio.NewReader(file)
	ar := make([]byte, 1024)
	for {
		i, _ := r.Read(ar)
		fmt.Println(string(ar[:i]))
		time.Sleep(time.Second)
	}

	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }
}
