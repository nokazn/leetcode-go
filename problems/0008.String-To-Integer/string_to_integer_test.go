package leetcode

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	testData := []struct {
		in  string
		out int
	}{
		{"42", 42},
		{"       -42", -42},
		{"4193 with words", 4193},
		{"   -4193 with words", -4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
		{"913137123122332", 2147483647},
		{"123.45", 123},
		{"  -123.", -123},
		{"  -123.9999", -123},
		{"  123.9999", 123},
		{"+-12", 0},
		{"-+12", 0},
		{"   +0 123", 0},
	}

	for i, tt := range testData {
		t.Run("string to integer "+strconv.Itoa(i), func(t *testing.T) {
			res := myAtoi(tt.in)
			ok := reflect.DeepEqual(res, tt.out)
			if ok {
				t.Logf("[input]: %v  [output]: %v", tt.in, res)
			} else {
				t.Errorf("[input]: %v  [output]: %v  [expected]: %v", tt.in, res, tt.out)
			}
		})
	}
}
