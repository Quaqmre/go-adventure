package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	a := os.TempDir()
	fmt.Println(a)

	tmpfile, _ := ioutil.TempFile("./", "qrcp")
	fmt.Println(tmpfile.Name())
	tmpfile.WriteString("selam")
	tmpfile.Close()
}
