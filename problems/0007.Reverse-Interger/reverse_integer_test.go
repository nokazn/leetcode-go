package leetcode

import (
	"strconv"
	"testing"
)

func TestReverse(t *testing.T) {
	testData := []struct {
		in  int
		out int
	}{
		{123, 321},
		{-123, -321},
		{-1230, -321},
		{120, 21},
		{0, 0},
		{1, 1},
		{1534236469, 0},
	}

	for i, tt := range testData {
		t.Run("reverse integer "+strconv.Itoa(i), func(t *testing.T) {
			res := reverse(tt.in)
			if res == tt.out {
				t.Logf("[input]: %v  [output]: %v", tt.in, res)
			} else {
				t.Errorf("[input]: %v  [output]: %v  [expected]: %v", tt.in, res, tt.out)
			}
		})
	}
}
