package lc

type ListNode struct {
	Val  int
	Next *ListNode
}

// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	next := &ListNode{}
	res := &ListNode{
		Val:  0,
		Next: next,
	}
	overflow := 0
	for l1 != nil || l2 != nil {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		v := v1 + v2 + overflow
		if v >= 10 {
			overflow = 1
			next.Val = v - 10
		} else {
			overflow = 0
			next.Val = v
		}
		if l1 != nil || l2 != nil {
			next.Next = &ListNode{}
			next = next.Next
		} else if overflow != 0 {
			next.Next = &ListNode{Val: overflow}
		}
	}

	return res.Next
}
