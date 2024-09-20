package lc

/*
	给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

	示例 1:
	输入: s = "abcabcbb"
	输出: 3

	示例 2:
	输入: s = "aab"
	输出: 2

	示例 3:
	输入: s = "pwwkew"
	输出: 3

	0 <= s.length <= 5 * 10^4
	s 由英文字母、数字、符号和空格组成
*/

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	result := 1
	startAt, cursor := 0, 0
	cache := map[byte]int{}
	for {
		if cursor == len(s) {
			return max(result, cursor-startAt)
		}
		cursorByte := s[cursor]
		if idx, found := cache[cursorByte]; found && idx >= startAt {
			result = max(result, cursor-startAt)
			startAt = idx + 1
		}
		cache[cursorByte] = cursor
		cursor++
	}
}

/*
给你一个字符串 s，找到 s 中最长的回文子串

示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

示例 2：
输入：s = "cbbd"
输出："bb"

提示：
1 <= s.length <= 1000
s 仅由数字和英文字母组成
*/
// https://leetcode.cn/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	return ""
}
