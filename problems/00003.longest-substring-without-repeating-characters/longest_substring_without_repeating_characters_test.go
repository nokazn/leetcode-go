package leetcode

import (
	"strconv"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	testData := []struct {
		in  string
		out int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"bbtablud", 6},
		{"bbtabalud", 5},
		{"bbtaalubd", 5},
		{"dvdf", 3},
		{"pwwkew", 3},
		{"", 0},
		{" ", 1},
		{"h", 1},
	}

	for i, tt := range testData {
		t.Run("longest substring without repeating characters "+strconv.Itoa(i), func(t *testing.T) {
			res := lengthOfLongestSubstring(tt.in)
			if res == tt.out {
				t.Logf("[input]: %v  [output]: %v", tt.in, res)
			} else {
				t.Errorf("[input]: %v  [output]: %v  [expected]: %v", tt.in, res, tt.out)
			}
		})
	}
}
