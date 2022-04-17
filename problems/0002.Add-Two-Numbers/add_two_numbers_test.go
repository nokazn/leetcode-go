package leetcode

import (
	"reflect"
	"testing"
)

type in struct {
	l1 []int
	l2 []int
}

type out []int

func intsToListNode(ints []int) *ListNode {
	if len(ints) == 0 {
		return nil
	}

	root := &ListNode{}
	current := root
	for _, v := range ints {
		current.Next = &ListNode{v, nil}
		current = current.Next
	}
	return root.Next
}

func listNodeToInts(l *ListNode) []int {
	var ints []int
	for l != nil {
		ints = append(ints, l.Val)
		l = l.Next
	}
	return ints
}

func TestAddTwoNumbers(t *testing.T) {
	testData := []struct {
		in
		out
	}{
		{
			in{[]int{2, 4, 3}, []int{5, 6, 4}},
			[]int{7, 0, 8},
		},
	}

	for _, tt := range testData {
		t.Run("add two numbers", func(t *testing.T) {
			res := listNodeToInts(addTwoNumbers(intsToListNode(tt.l1), intsToListNode(tt.l2)))
			// 型を揃えないと比較できない
			ok := reflect.DeepEqual(res, []int(tt.out))
			if ok {
				t.Logf("[input]: %v  [output]: %v", tt.in, res)
			} else {
				t.Errorf("[input]: %v  [output]: %v  [expected]: %v", tt.in, res, tt.out)
			}
		})
	}
}
