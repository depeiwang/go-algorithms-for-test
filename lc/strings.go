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

/*
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s 的，而不是部分字符串。

示例 1：
输入：s = "aa", p = "a"
输出：false
解释："a" 无法匹配 "aa" 整个字符串。

示例 2:
输入：s = "aa", p = "a*"
输出：true
解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。

示例 3：
输入：s = "ab", p = ".*"
输出：true
解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。

提示：
1 <= s.length <= 20
1 <= p.length <= 20
s 只包含从 a-z 的小写字母。
p 只包含从 a-z 的小写字母，以及字符 . 和 *。
保证每次出现字符 * 时，前面都匹配到有效的字符
*/
// https://leetcode.cn/problems/regular-expression-matching/
func isMatch(s string, p string) bool {
	// TODO
	sl := len(s)
	j := 0
	lastChar := byte(0)
	match := true
	for i := 0; i < sl; i++ {
		pc := p[j]
		if pc != '.' && pc != '*' {
			if pc != s[i] {
				match = false
				break
			}
			lastChar = pc
			j++
			continue
		}
		if pc == '.' {
			lastChar = pc
			j++
			continue
		}
		if pc == '*' {

		}
	}

	return match && lastChar != 0
}
