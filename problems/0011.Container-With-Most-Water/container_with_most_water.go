package leetcode

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxArea(height []int) int {
	volume := 0
	length := len(height)
	start := 0
	end := length - 1
	for start < end {
		h1 := height[start]
		h2 := height[end]
		volume = max(volume, (end-start)*min(h1, h2))
		if h1 > h2 {
			end -= 1
		} else {
			start += 1
		}
	}
	return volume
}
