// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 228.

// Pipeline1 demonstrates an infinite 3-stage pipeline.
package main

import "fmt"

//!+
// 该示例程序由三个goroutine, 两个通道构成
// Counter goroutine 负责生产数字0, 1, 2..., 并将生产的数字发送到阻塞通道naturals
// Squarer goroutine 负责从阻塞通道naturals中取出数字, 计算对应平方, 并将计算结果发送到阻塞通道squares
// Main goroutine 负责从阻塞通道squares中取出计算结果并打印
func main() {
	// 新建阻塞通道naturals, squares
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	// 通过go关键字创建新的goroutine, 并由后台调度执行
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}

//!-
