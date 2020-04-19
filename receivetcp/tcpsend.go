package main

import (
	"fmt"
	"net"
)

func main() {
	lncock, _ := net.Listen("tcp", ":9001")

	for {

		conn, _ := lncock.Accept()

		recvBuffer := make([]byte, 256)

		i, _ := conn.Read(recvBuffer)

		str := string(recvBuffer[:i])

		fmt.Println(str)
	}

}
