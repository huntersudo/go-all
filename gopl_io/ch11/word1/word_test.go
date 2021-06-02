// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+test
package word

import "testing"

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
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
func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") = false`)
	}
}

// 发现第⼀个BUG的原因是我们采⽤了 byte⽽不是rune序列，所以 像“été”中的é等⾮ASCII字符不能正确处理
/**
=== RUN   TestFrenchPalindrome
    word_test.go:32: IsPalindrome("été") = false
--- FAIL: TestFrenchPalindrome (0.00s)

FAIL

*/

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}

/**
=== RUN   TestCanalPalindrome
    word_test.go:45: IsPalindrome("A man, a plan, a canal: Panama") = false
--- FAIL: TestCanalPalindrome (0.00s)

FAIL
*/
//!-more
