package transformer

type ListNode struct {
	Val  int
	Next *ListNode
}

func IntsToListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	root := &ListNode{}
	current := root
	for _, v := range nums {
		current.Next = &ListNode{v, nil}
		current = current.Next
	}
	return root.Next
}
