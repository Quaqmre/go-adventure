package main

import (
	"fmt"
	"os"
)

func main() {
	// viper.SetDefault("asd", "123")
	t := os.Getenv("asd")
	fmt.Println(t)
}
