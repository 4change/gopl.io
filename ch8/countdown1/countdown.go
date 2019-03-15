// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 244.

// Countdown implements the countdown for a rocket launch.
package main

import (
	"fmt"
	"time"
)

//!+
// 火箭发射倒计时程序
func main() {
	fmt.Println("Commencing countdown.")
	// time.Tick()函数返回一个通道, 该通道定期发送事件, 事件值为一个时间戳
	tick := time.Tick(1 * time.Second)
	// 火箭发射的10s倒计时
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		// 从tick通道中取出时间戳,该时间戳每秒通过事件发送到通道
		<-tick
	}
	launch()
}

//!-

func launch() {
	fmt.Println("Lift off!")
}
