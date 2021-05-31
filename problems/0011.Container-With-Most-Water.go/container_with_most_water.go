package leetcode

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxArea(height []int) int {
	var volume int
	l := len(height)
	for i1, h1 := range height {
		for i2 := i1 + 1; i2 < l; i2 += 1 {
			h2 := height[i2]
			v := min(h1, h2) * (i2 - i1)
			volume = max(volume, v)
		}
	}
	return volume
}
