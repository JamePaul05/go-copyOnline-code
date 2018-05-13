package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func receiveFile(fileName string, coon net.Conn) {
	file, createErr := os.Create(fileName)
	if createErr != nil {
		fmt.Println("createErr's error is ", createErr)
		return
	}
	readBuf := make([]byte, 1024)
	for {
		n, readErr := coon.Read(readBuf)
		if readErr != nil {
			if readErr == io.EOF {
				fmt.Println("文件读取完了")
				return
			} else {
				fmt.Println("readErr's error is ", readErr)
				return
			}
		}
		file.Write(readBuf[:n])
	}
	defer coon.Close()

}
func main() {

	listener, listenErr := net.Listen("tcp", "127.0.0.1:8000") //对服务器进行监听
	if listenErr != nil {
		fmt.Println("listenErr's error is ", listenErr)
		return
	}
	defer listener.Close()

	coon, acceotErr := listener.Accept() //接受信息
	if acceotErr != nil {
		fmt.Println("acceotErr's error is ", acceotErr)
		return
	}
	defer coon.Close()
	readBuf := make([]byte, 1024) //读取写过来的信息  读取成功了写入 OK
	n, readErr := coon.Read(readBuf)
	if readErr != nil {
		fmt.Println("readErr's error is ", readErr)
		return
	}
	fileName := string(readBuf[:n])
	_, err := coon.Write([]byte("ok"))
	if err != nil {
		fmt.Println("err", err)
		return
	}

	receiveFile(fileName, coon)

}
