// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 222.

// Clock is a TCP server that periodically writes the time.
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

// 并发时钟服务器
// 主goroutine接收客户端连接请求, 并对每一个到来的客户端请求创建一个goroutine进行处理
func main() {
	// 本地服务器通过TCP协议连接8000端口
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		// 本地服务器接收请求
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		// 通过go关键字创建协程, 使当前服务器支持协程, 每到来一个客户端连接便创建一个协程进行处理
		go handleConn(conn) // handle connections concurrently
	}
	//!-
}
