package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// doClean 模拟做程序的清理工作
func doClean(closed chan struct{}) {
	// time.Sleep 模拟清理工作
	time.Sleep(60 * time.Second)
	// 清理完成以后关闭closed channel
	close(closed)
}

func main() {
	var closing = make(chan struct{})
	var closed = make(chan struct{})

	go func() {
		for {
			select {
			case <-closing:
				return
			default:
				// 模拟业务处理
				time.Sleep(1000 * time.Second)
			}
		}
	}()

	notifyC := make(chan os.Signal)
	signal.Notify(notifyC, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("等待程序退出信号")
	<-notifyC

	// 关闭执行执行的业务程序
	close(closing)

	// 开始执行清理工作，清理完成以后会返回closed的channel
	go doClean(closed)

	// 防止清理工作执行太久，执行超时操作
	select {
	case <-closed:
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
	fmt.Println("process is exit")

}
