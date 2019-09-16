// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+test
package word

import "testing"

/*
Go中功能测试函数必须以Test作为前缀
*/
func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		// t.Error()方法, 输出测试中的错误信息
		t.Error(`IsPalindrome("detartrated") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}

//!-test

// The tests below are expected to fail.
// See package gopl.io/ch11/word2 for the fix.

//!+more
// 法语环境回文测试
func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") = false`)
	}
}

// 南美环境回文测试
func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}

//!-more
