package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

//向服务器的统计目录复制指定的文件
func copyFile(path string, coon net.Conn) {
	file, openErr := os.Open(path) //以只读的方式打开文件
	if openErr != nil {
		fmt.Println("openErr's error is ", openErr)
		return
	}
	defer file.Close()
	readBuf := make([]byte, 1024*5)
	for {
		n, readErr := file.Read(readBuf) //读取文件内容
		if readErr != nil {
			if readErr == io.EOF {
				fmt.Println("文件已经读取完毕")
			} else {
				fmt.Println("openErr's error is ", readErr)
			}
			return
		}
		coon.Write(readBuf[:n]) //读取的是什么就写入什么

	}

}

func main() {
	var path string
	fmt.Println("请输入需要拷贝的文件")

	fmt.Scan(&path)                    //获取命令行的输入的文件名
	fileInfo, fileErr := os.Stat(path) //这个函数返回的是文件
	if fileErr != nil {
		fmt.Println(" fileErr's error is  ", fileErr)
		return
	}

	coon, dialErr := net.Dial("tcp", "127.0.0.1:8000")
	if dialErr != nil {
		fmt.Println(" dialErr's error is  ", dialErr)
		return
	}
	defer coon.Close() //在函数结束的一瞬间关闭

	_, writeErr := coon.Write([]byte(fileInfo.Name()))
	if writeErr != nil {
		fmt.Println(" writeErr's error is  ", writeErr)
		return
	}
	readBuf := make([]byte, 1024)
	n1, readErr := coon.Read(readBuf) //从键盘读取了多少就写多少
	if readErr != nil {
		fmt.Println(" writeErr's error is  ", readErr)
		return
	}
	if "ok" == string(readBuf[:n1]) {
		copyFile(path, coon)

	}

}
