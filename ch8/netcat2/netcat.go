// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 223.

// Netcat is a simple read/write client for TCP servers.
package main

import (
	"io"
	"log"
	"net"
	"os"
)

//!+
func main() {
	// 客户端发起拨号连接, 通过TCP协议连接到本地8000端口
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	// 延迟函数,进行连接的关闭
	defer conn.Close()
	// 新起goroutine, 将连接数据复制到Stdout, 以进行输出
	go mustCopy(os.Stdout, conn)
	// 主goroutine从Stdin获取输入, 并发送到连接conn中
	mustCopy(conn, os.Stdin)
}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
