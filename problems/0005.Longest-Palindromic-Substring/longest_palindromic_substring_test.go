package leetcode

import (
	"strconv"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	testData := []struct {
		in  string
		out string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"cbabsbad", "absba"},
		{"a", "a"},
		{"aaaaa", "aaaaa"},
		{"acaa", "aca"},
		{"aca", "aca"},
		{"acaa", "aca"},
		{"ac", "a"},
		{"", ""},
	}

	for i, tt := range testData {
		t.Run("reverse integer "+strconv.Itoa(i), func(t *testing.T) {
			res := longestPalindrome(tt.in)
			if res == tt.out {
				t.Logf("[input]: %v  [output]: %v", tt.in, res)
			} else {
				t.Errorf("[input]: %v  [output]: %v  [expected]: %v", tt.in, res, tt.out)
			}
		})
	}
}
