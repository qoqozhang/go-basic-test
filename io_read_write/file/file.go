package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

/*
文件源 --> io.Reader --> []byte --> io.Writer --> 文件
*/
func main() {
	file, err := os.Open("./read.txt")
	wFile, err := os.OpenFile("./write.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()
	defer wFile.Close()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 1)
	for {
		file.Seek(1, 1) //读取的游标移动,相对当前位置位移一个
		rn, err := file.Read(buf)
		if err == io.EOF {
			log.Fatal("文件读完了")
			break
		}
		fmt.Printf("read: %v\n", string(buf[:rn]))
		//wn, _ := wFile.Write(buf[:rn])
		//fmt.Printf("write %d 字节\n", wn)
	}
}
