package leetcode

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func quotientWithRemainder(dividend, divisor int) int {
	q := dividend / divisor
	// 余りがある場合加算
	if dividend%divisor > 0 {
		return q + 1
	}
	return q
}

func convert(s string, numRows int) string {
	// 最小値は 1
	lengthInGroups := max((numRows-1)*2, 1)
	l := len(s)
	numGroups := quotientWithRemainder(l, lengthInGroups)

	var bytes []byte
	for m := 0; m < lengthInGroups/2+1; m++ {
		for n := 0; n < numGroups; n++ {
			i1 := lengthInGroups*n + m
			if i1 >= l {
				continue
			}

			i2 := lengthInGroups*n + lengthInGroups - m
			if m > 0 && i1 < i2 && i2 < l {
				bytes = append(bytes, s[i1], s[i2])
			} else {
				bytes = append(bytes, s[i1])
			}
		}
	}

	return string(bytes)
}
