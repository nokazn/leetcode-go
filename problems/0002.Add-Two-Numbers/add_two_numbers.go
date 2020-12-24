package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{0, nil}
	current := root
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		val := carry
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}

		// 現在の current の Next
		current.Next = &ListNode{val % 10, nil}
		// 次の current と carry
		current = current.Next
		carry = val / 10
	}

	return root.Next
}
