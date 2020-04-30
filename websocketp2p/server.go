package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		}}

	handler := func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer func() {
			err := conn.Close()
			if err != nil {
				log.Println("Error:", err.Error())
			}
		}()

		log.Println("Added new client. Now", "clients connected.")
		for {
			messageType, reader, err := conn.NextReader()
			time.Sleep(time.Second * 2)
			x, _ := ioutil.ReadAll(reader)
			fmt.Println(string(x))
			_ = reader
			if err != nil {
				log.Println(err)
				return
			} else if messageType != websocket.BinaryMessage {
				log.Println("Non binary message recived, ignoring")
			}

			// _ = conn.WriteMessage(websocket.BinaryMessage, data)
		}
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe("192.168.1.38:4000", nil)
}
