// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

// 存钱通道
var deposits = make(chan int) // send amount to deposit
// 余额查询通道
var balances = make(chan int) // receive balance

func Deposit(amount int) {
	// 向存钱通道存入
	deposits <- amount
}

func Balance() int {
	// 从余额查询通道读出
	return <-balances
}

func teller() {
	// balance变量由监控协程teller进行修改
	var balance int // balance is confined to teller goroutine
	for {
		select {
		// 从存钱通道取出金额, 放入balance变量
		case amount := <-deposits:
			balance += amount
		// 从balance变量取出余额, 并放入余额查询通道
		case balances <- balance:
		}
	}
}

// 包初始化时执行teller监控goroutine
func init() {
	go teller() // start the monitor goroutine
}

//!-
