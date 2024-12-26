package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buff := make([]byte, 1024)

	bytes_received, err := conn.Read(buff)

	if err != nil {
		fmt.Println("Error occured: ", err)
		return
	}
	fmt.Println(string(buff))
	fmt.Println("bytes recv", bytes_received)
	bytes_sent, err := conn.Write(buff)

	if err != nil {
		fmt.Println("Error occured: ", err)
		return
	}

	fmt.Println("bytes sent", bytes_sent)
}
