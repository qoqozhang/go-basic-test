package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// doClean 模拟做程序的清理工作
func doClean(closed chan struct{}) {
	// time.Sleep 模拟清理工作
	i := rand.Int() % 10
	fmt.Printf("开始执行%d 秒的清理工作.\n", i)
	time.Sleep(time.Second * time.Duration(i))
	// 清理完成以后关闭closed channel
	close(closed)
}

func main() {
	//TODO 模拟程序正常工作
	fmt.Println("程序运行中...")

	// 捕获程序退出信号
	var closed = make(chan struct{})
	notifyC := make(chan os.Signal)
	signal.Notify(notifyC, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("键入CTRL + C 结束程序...")
	<-notifyC
	fmt.Println("开始执行清理工作...")

	// 开始执行清理工作，清理完成以后会返回closed的channel
	go doClean(closed)

	// 防止清理工作执行太久，执行超时操作
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	select {
	case <-closed:
		fmt.Println("清理完成...")
	case <-ctx.Done():
		fmt.Println("清理工作超时...")
	}
	fmt.Println("程序退出完成...")
}
