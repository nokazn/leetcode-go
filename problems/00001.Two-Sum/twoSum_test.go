package leetcode

import (
	"strconv"
	"testing"
)

type in struct {
	nums   []int
	target int
}

type out = []int

func TestTwoSum(t *testing.T) {
	testData := []struct {
		in  in
		out out
	}{
		{
			in{[]int{2, 7, 11, 15}, 9},
			[]int{0, 1},
		},
		{
			in{[]int{3, 2, 4}, 6},
			[]int{1, 2},
		},
		{
			in{[]int{3, 3, 6}, 6},
			[]int{0, 1},
		},
		{
			in{[]int{3, 4, 8}, 6},
			[]int{0, 0},
		},
		{
			in{[]int{3, 2, 3, 6}, 8},
			[]int{1, 3},
		},
	}

	for i, tt := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res := twoSum(tt.in.nums, tt.in.target)
			if tt.out[0] == res[0] && tt.out[1] == res[1] {
				t.Logf("[input]: %v  [output]: %v\n", tt.in, res)
			} else {
				t.Errorf("[input]: %v  [output]: %v  [expected]: %v", tt.in, res, tt.out)
			}
			t.Log("\n")
		})
	}
}
