// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 229.

// Pipeline2 demonstrates a finite 3-stage pipeline.
package main

import "fmt"

//!+
// 该示例程序由三个goroutine, 两个通道构成
// Counter goroutine 负责生产数字0, 1, 2..., 并将生产的数字发送到阻塞通道naturals, 在数据发送完毕后通过close()函数关闭通道
// Squarer goroutine 负责从阻塞通道naturals中取出数字, 计算对应平方, 并将计算结果发送到阻塞通道squares, 在数据发送完毕后通过close()函数关闭通道
// Main goroutine 负责从阻塞通道squares中取出计算结果并打印
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		// 通过close()函数关闭通道
		// 在通道关闭后, 便无法再向通道中发送数据, 但可以从通道中读取数据, 此时从通道中读出的数据为对应通道类型的零值
		// range循环支持从通道中接收数据, 它会判断从通道中读出的数据是否为通道的最后一个有效数据, 在读出通道中最后一个有效数据后会自动结束循环
		close(naturals)
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}

//!-
