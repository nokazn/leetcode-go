package leetcode

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxArea(height []int) int {
	volume := 0
	for start, h1 := range height {
		for x, h2 := range height[start+1:] {
			s := (x + 1) * min(h1, h2)
			if s > volume {
				volume = s
			}
		}
	}
	return volume
}
