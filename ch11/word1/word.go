// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 303.
//!+

// Package word provides utilities for word games.
package word

// IsPalindrome reports whether s reads the same forward and backward.
// (Our first attempt.)
// IsPalindrome()函数,用于判断一个字符串是否为回文字符串
// bug1: 回文串使用字节序列而不是字符序列进行比较, 会出现程序异常
// bug2: 未忽略空格、标点符号和字母大小写
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

//!-
