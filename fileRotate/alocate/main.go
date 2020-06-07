package main

func main() {
	// // Buffera almıyor 1024 1024 okuyor
	// f, _ := os.OpenFile("./test.img", os.O_RDWR, 0666)
	// defer f.Close()
	// f2, _ := os.OpenFile("testf2", os.O_CREATE, 0666)
	// defer f2.Close()
	// by := make([]byte, 1024)
	// for {

	// 	_, err := f.Read(by)
	// 	if err == io.EOF {
	// 		return
	// 	}
	// 	f2.Write(by)
	// }

	// // tüm 1 gb lık instanceyi buffera alıyor
	// a, _ := ioutil.ReadFile("test.img")
	// _ = a
	// for {

	// }
}
