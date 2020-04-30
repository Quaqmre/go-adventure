package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	u := url.URL{Scheme: "ws", Host: "192.168.1.38" + ":4000", Path: "/"}
	fmt.Println(u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	botName := fmt.Sprintf("%s %v", "test", 4)

	time.Sleep(time.Second * 3)
	c.WriteMessage(websocket.BinaryMessage, []byte(botName))
	c.WriteMessage(websocket.BinaryMessage, []byte("selambebek"))
}
