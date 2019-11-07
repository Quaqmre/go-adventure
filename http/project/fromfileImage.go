package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Get(w http.ResponseWriter, r *http.Request) {
	imgFile, err := os.Open("ObjectStorageError.png") // a QR code image

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer imgFile.Close()

	// create a new buffer base on file size
	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	// if you create a new image instead of loading from file, encode the image to buffer instead with png.Encode()

	// png.Encode(&buf, image)

	// convert the buffer bytes to base64 string - use buf.Bytes() for new image
	imgBase64Str := base64.StdEncoding.EncodeToString(buf)

	// Embed into an html without PNG file
	img2html := fmt.Sprintf(`{"alive":"%s"}`, imgBase64Str)

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, img2html)

}

type PostModel struct {
	Alive string `json:"alive"`
	Name  string `json:"name"`
}

func Post(w http.ResponseWriter, r *http.Request) {
	var post PostModel
	_ = json.NewDecoder(r.Body).Decode(&post)

	a, _ := base64.StdEncoding.DecodeString(post.Alive)
	img, _, _ := image.Decode(bytes.NewReader(a))

	f, _ := os.Create("h.png")
	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
	_ = a
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Get).Methods("GET")
	router.HandleFunc("/post", Post).Methods("POST")

	http.ListenAndServe(":8000", router)

}
