package leetcode

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	p := make([]string, n)
	for i := range p {
		j := i + 1
		if j%15 == 0 {
			p[i] = "FizzBuzz"
		} else if j%5 == 0 {
			p[i] = "Buzz"
		} else if j%3 == 0 {
			p[i] = "Fizz"
		} else {
			p[i] = strconv.Itoa(j)
		}
	}
	return p
}
