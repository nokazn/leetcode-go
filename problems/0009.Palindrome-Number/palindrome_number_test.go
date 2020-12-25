package leetcode

import (
	"reflect"
	"strconv"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testData := []struct {
		in  int
		out bool
	}{
		{121, true},
		{-121, false},
		{3, true},
		{-3, false},
		{149080941, true},
		{14900941, true},
		{149009141, false},
		{10, false},
		{-101, false},
		{101, false},
	}

	for i, tt := range testData {
		t.Run("palindrome number "+strconv.Itoa(i), func(t *testing.T) {
			res := isPalindrome(tt.in)
			ok := reflect.DeepEqual(res, tt.out)
			if ok {
				t.Logf("[input]: %v  [output]: %v", tt.in, res)
			} else {
				t.Errorf("[input]: %v  [output]: %v  [expected]: %v", tt.in, res, tt.out)
			}
		})
	}
}
