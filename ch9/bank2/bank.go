// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 262.

// Package bank provides a concurrency-safe bank with one account.
package bank

//!+
var (
	// 互斥信号量
	sema    = make(chan struct{}, 1) // a binary semaphore guarding balance
	// 由互斥信号量控制的变量
	balance int
)

func Deposit(amount int) {
	// 获取互斥信号量,相当于获取锁
	sema <- struct{}{} // acquire token
	balance = balance + amount
	// 释放互斥信号量,相当于释放锁
	<-sema // release token
}

func Balance() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // release token
	return b
}

//!-
