package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

var (
	port = flag.String("port", "9000", "port used for ws connection")
)

func main() {
	flag.Parse()
	conn, _ := net.Dial("tcp", "localhost:9001/hijack")
	log.Println("Incoming client connection")
	go func() {
		for {
			recvBuffer := make([]byte, 256)
			bytesRead, err := conn.Read(recvBuffer)
			if err != nil {
				return
			}
			// t := string(m)
			data := recvBuffer[:bytesRead]
			g := string(data)
			fmt.Println(g)
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		a := []byte(text)
		_, err := conn.Write(a)
		if err != nil {
			fmt.Println("Error sending message: ", err.Error())
			break
		}
	}

}
