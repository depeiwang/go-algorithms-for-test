package archive

// https://leetcode.cn/problems/two-sum/description/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}

// https://leetcode.cn/problems/add-two-numbers/description/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	cur := root
	third := 0
	for l1 != nil || l2 != nil || third != 0 {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum := v1 + v2 + third
		third = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}
	return root.Next
}

// https://leetcode.cn/problems/string-to-integer-atoi/
func myAtoi(s string) int {
	num := 0
	sign := 1
	start := false
	const max = 1<<31 - 1
	for _, c := range s {
		if c == ' ' {
			if start {
				break
			}
			continue
		}
		if c == '-' {
			if start {
				break
			}
			start = true
			sign = -1
			continue
		}
		if c == '+' {
			if start {
				break
			}
			start = true
			continue
		}
		if c < '0' || c > '9' {
			break
		}
		start = true

		digit := int(c - '0')
		if num > max/10 || (num == max/10 && digit > 7) {
			if sign == 1 {
				return max
			}
			return -1 << 31
		}
		num = num*10 + digit
	}
	return num * sign
}

/*
n=5, d=2n-2=8
0     8         16  --> rowNum=0 [a0=rowNum]
1   7 9      15 17  --> rowNum=1 [a0=rowNum, a0=d-rowNum]
2  6  10   14   18  --> rowNum=2 [a0=rowNum, a0=d-rowNum]
3 5   11 13     19  --> rowNum=3 [a0=rowNum, a0=d-rowNum]
4     12        20  --> rowNum=4 [a0=rowNum]
*/
// https://leetcode.cn/problems/zigzag-conversion/description/
func convert(s string, numRows int) string {
	totalLen := len(s)
	if numRows == 1 || totalLen <= numRows {
		return s
	}

	resultIdx := 0
	result := make([]byte, totalLen)

	for rowNum := 0; rowNum < numRows; rowNum++ {
		d := (numRows << 1) - 2
		a01 := rowNum
		a02 := d - rowNum
		for {
			if a01 >= totalLen {
				break
			}

			result[resultIdx] = s[a01]
			resultIdx++
			a01 += d

			if rowNum > 0 && rowNum < numRows-1 && a02 < totalLen {
				result[resultIdx] = s[a02]
				resultIdx++
				a02 += d
			}

		}
	}
	return string(result)
}

// https://leetcode.cn/problems/reverse-integer/description/
func reverse(x int) int {
	const MIN = (-1 << 31) / 10
	const MAX = ((1 << 31) - 1) / 10
	var res int

	for x != 0 {
		digit := x % 10
		if res > MAX || (res == MAX && digit > 7) {
			return 0
		}
		if res < MIN || (res == MIN && digit < -8) {
			return 0
		}
		res = res*10 + x%10
		x /= 10
	}

	return res
}

// https://leetcode.cn/problems/palindrome-number/description/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	origin := x
	const MAX = ((1 << 31) - 1) / 10
	reverse := 0

	for x > 0 {
		if reverse > MAX {
			return false
		}
		reverse = reverse*10 + x%10
		x /= 10
	}
	return reverse == origin
}

// https://leetcode.cn/problems/container-with-most-water/
func maxArea(height []int) int {

	max := 0
	for i, j := 0, len(height)-1; i < j; {
		area := min(height[i], height[j]) * (j - i)
		if area > max {
			max = area
		}
		// 在长度缩小的情况下，如果高度不变或者变小，那么面积一定变小
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

// https://leetcode.cn/problems/integer-to-roman/
func intToRoman(num int) string {
	mapping := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	result := ""
	for {
		if num == 0 {
			break
		}
		if num >= 1000 {
			result += mapping[1000]
			num -= 1000
		} else if num >= 900 {
			result += mapping[900]
			num -= 900
		} else if num >= 500 {
			result += mapping[500]
			num -= 500
		} else if num >= 400 {
			result += mapping[400]
			num -= 400
		} else if num >= 100 {
			result += mapping[100]
			num -= 100
		} else if num >= 90 {
			result += mapping[90]
			num -= 90
		} else if num >= 50 {
			result += mapping[50]
			num -= 50
		} else if num >= 40 {
			result += mapping[40]
			num -= 40
		} else if num >= 10 {
			result += mapping[10]
			num -= 10
		} else if num >= 9 {
			result += mapping[9]
			num -= 9
		} else if num >= 5 {
			result += mapping[5]
			num -= 5
		} else if num >= 4 {
			result += mapping[4]
			num -= 4
		} else {
			result += mapping[1]
			num--
		}
	}

	return result
}

// https://leetcode.cn/problems/roman-to-integer/
func romanToInt(s string) int {
	mapping := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	result := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			if v, ok := mapping[s[i:i+2]]; ok {
				result += v
				i++
				continue
			}
		}
		result += mapping[s[i:i+1]]
	}
	return result
}

// https://leetcode.cn/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	minLen := len(strs[0])
	for _, str := range strs[1:] {
		if len(str) < minLen {
			minLen = len(str)
		}
	}
	if minLen == 0 {
		return ""
	}

	for i := 0; i < minLen; i++ {
		c := strs[0][i]
		for _, str := range strs[1:] {
			if str[i] != c {
				return str[:i]
			}
		}
	}
	return strs[0][:minLen]
}

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	mapping := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	result := []string{""}
	for i := 0; i < len(digits); i++ {
		digit := digits[i]
		letters, ok := mapping[digit]
		if !ok {
			continue
		}
		temp := make([]string, 0, len(result)*len(letters))
		for _, prefix := range result {
			for j := 0; j < len(letters); j++ {
				temp = append(temp, prefix+string(letters[j]))
			}
		}
		result = temp
	}
	return result
}
