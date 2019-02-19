// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 305.
//!+

// Package word provides utilities for word games.
// 包word提供了文字游戏相关的工具函数
package word

import "unicode"

// IsPalindrome reports whether s reads the same forward and backward.
// Letter case is ignored, as are non-letters.
// IsPalindrome判断一个字符串是否是回文字符串
// 忽略字母大小写,以及非字母字符
func IsPalindrome(s string) bool {
	var letters []rune
	// 提付字符串s中的字母字符
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

//!-
