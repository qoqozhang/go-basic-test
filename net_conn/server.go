package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func Server() {
	fmt.Println("Server begin running...")
	l, err := net.Listen("tcp", "10.1.1.114:2000")
	if err != nil {
		log.Fatalf("server start faild: %v\n", err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("client connect failed: %v\n", err)
		}
		go func(c net.Conn) {
			io.Copy(c, c)
			c.Close()
		}(conn)
	}
}
