// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 219.
//!+

// Clock1 is a TCP server that periodically writes the time.
package main

import (
	"io"
	"log"
	"net"
	"time"
)

// 顺序时钟服务器程序
func main() {
	// 开启一个本地8000端口的TCP连接
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		// 本地服务器接受请求
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		// 本地服务器开始处理连接
		handleConn(conn) // handle one connection at a time
	}
}

func handleConn(c net.Conn) {
	// 延迟函数,在程序退出时执行网络连接Conn的关闭
	defer c.Close()
	for {
		// 格式化当前时间并输出当网络连接Conn中
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		// 休眠1s, 重新开始循环
		time.Sleep(1 * time.Second)
	}
}

//!-
