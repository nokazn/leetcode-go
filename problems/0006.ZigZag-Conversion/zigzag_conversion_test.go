package leetcode

import (
	"strconv"
	"testing"
)

type in struct {
	s    string
	rows int
}

func TestConvert(t *testing.T) {
	testData := []struct {
		in
		out string
	}{
		{in{"PAYPALISHIRING", 4}, "PINALSIGYAHRPI"},
		{in{"A", 1}, "A"},
	}

	for i, tt := range testData {
		t.Run("zigzag conversion "+strconv.Itoa(i), func(t *testing.T) {
			res := convert(tt.in.s, tt.in.rows)
			if res == tt.out {
				t.Logf("[input]: %v  [output]: %v", tt.in, res)
			} else {
				t.Errorf("[input]: %v  [output]: %v  [expected]: %v", tt.in, res, tt.out)
			}
		})
	}
}
