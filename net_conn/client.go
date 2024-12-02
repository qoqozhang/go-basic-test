package main

import (
	"fmt"
	"log"
	"net"
)

func Client() {
	fmt.Println("client connect to server...")
	conn, err := net.Dial("tcp", "10.1.1.236:2000")
	if err != nil {
		log.Fatalf("connect to server faild: %v\n", err)
	}
	conn.Write([]byte("zhang"))
	defer conn.Close()
}
