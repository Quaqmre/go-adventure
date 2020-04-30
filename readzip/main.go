package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

// func main() {
// 	f, _ := os.OpenFile("./xml.zip", os.O_RDONLY, 0755)
// 	by, _ := ioutil.ReadAll(f)
// 	a := string(by)
// 	arr := []byte{}
// 	for _, i := range a {
// 		bt := []byte{}
// 		_ = utf8.EncodeRune(bt, i)
// 		arr = append(arr, bt...)
// 	}
// 	fmt.Println(string(arr))
// 	defer f.Close()
// }

func main() {
	f, _ := os.OpenFile("./xmlput.zip", os.O_RDONLY, 0755)
	all, _ := ioutil.ReadAll(f)
	encoded := base64.StdEncoding.EncodeToString(all)

	fmt.Println(encoded)
}
