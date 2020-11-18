package leetcode

import (
	"reflect"
	"strconv"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	testData := []struct {
		in  int
		out []string
	}{
		{
			15,
			[]string{
				"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
				"Fizz",
				"7",
				"8",
				"Fizz",
				"Buzz",
				"11",
				"Fizz",
				"13",
				"14",
				"FizzBuzz",
			},
		},
	}

	for i, tt := range testData {
		t.Run("fizz buzz "+strconv.Itoa(i), func(t *testing.T) {
			res := fizzBuzz(tt.in)
			ok := reflect.DeepEqual(res, tt.out)
			if ok {
				t.Logf("[input]: %v  [output]: %v", tt.in, res)
			} else {
				t.Errorf("[input]: %v  [output]: %v  [expected]: %v", tt.in, res, tt.out)
			}
		})
	}
}
