package main

import "sync"

// func main() {
// dir := http.Dir("./")
// go http.ListenAndServe(":8080", http.FileServer(dir)) // file browserda gösterilerek içerik görüntülenebilmektedir.

// http.HandleFunc("/readme", readme)
// go http.ListenAndServe(":8081", nil) // go olmasının nedeni concurrent olarak serverları başlatabilmektir.

// time.Sleep(10 * time.Second)
// }

// func readme(res http.ResponseWriter, req *http.Request) {
// 	http.ServeFile(res, req, "./jsonParse")
// }

//----

func main() {
	var mutex sync.RWMutex
	mutex.RLock()
}
