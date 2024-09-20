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
