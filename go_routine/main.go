package main

import (
	"fmt"
	"time"
)

func main() {
	output := make(chan string, 10)
	go write(output)
	for s := range output {
		fmt.Println("res: ", s)
		time.Sleep(time.Second)
	}
}

func write(ch chan string) {
	for {
		select {
		case ch <- "hello":
			fmt.Println("write hello")
		default:
			fmt.Println("channel full")
		}
		time.Sleep(time.Millisecond * 500)
	}
}
