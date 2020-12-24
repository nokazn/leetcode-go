package leetcode

import (
	"reflect"
	"testing"

	"github.com/nokazn/leetcode-go/transformer"
)

type in struct {
	l1 []int
	l2 []int
}

type out []int

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
			res := addTwoNumbers(transformer.IntsToListNode(tt.l1), transformer.IntsToListNode(tt.l2))
			ok := reflect.DeepEqual(res, tt.out)
			if ok {
				t.Logf("[input]: %v  [output]: %v", tt.in, res)
			} else {
				t.Errorf("[input]: %v  [output]: %v  [expected]: %v", tt.in, res, tt.out)
			}
		})
	}
}
