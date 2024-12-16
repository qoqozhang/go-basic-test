package main

import (
	"fmt"
	"testing"
)

func recv(c chan int) {
	ret := <-c
	fmt.Println("receive success: ", ret)
}

func TestMain(m *testing.M) {
	ch := make(chan int)
	go recv(ch)
	ch <- 10
	fmt.Println("send success")
}
