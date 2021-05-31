package leetcode

import (
	"reflect"
	"strconv"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testData := []struct {
		in  []int
		out int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	}

	for i, tt := range testData {
		t.Run("container with most water"+strconv.Itoa(i), func(t *testing.T) {
			res := maxArea(tt.in)
			ok := reflect.DeepEqual(res, tt.out)
			if ok {
				t.Logf("[input]: %v  [output]: %v", tt.in, res)
			} else {
				t.Errorf("[input]: %v  [output]: %v  [expected]: %v", tt.in, res, tt.out)
			}
		})
	}
}
