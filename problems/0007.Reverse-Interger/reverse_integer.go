package leetcode

import (
	"math"
)

func reverse(x int) int {
	var res int
	for x != 0 {
		p := x % 10
		x /= 10
		res = res*10 + p
		if res < math.MinInt32 || math.MaxInt32 < res {
			return 0
		}
	}

	return res
}
