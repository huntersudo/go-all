package main

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	// 查找以s[0]为开头的最长回文子串
	right := len(s)
	for ; right >= 0; right-- {
		// 找到最长的回文子串就退出
		if isPalindrome(s[:right]) {
			break
		}
	}

	// 用递归的方式获取s[1:]的长度，与上面的结果对比
	subS := longestPalindrome(s[1:])
	if len(subS) > right {
		return subS
	}
	return s[:right]
}

// 头尾两个指针不断往中间缩进
func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	left, right := 0, len(s) - 1

	// 如果s长度为奇数，退出时刚好left=right，偶数则为 left > right
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
