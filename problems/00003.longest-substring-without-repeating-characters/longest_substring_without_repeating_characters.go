package leetcode

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func lengthOfLongestSubstring(s string) int {
	set := make(map[rune]int)
	startAt, maxCounts, interval := 0, 0, 0
	for i, v := range s {
		prev, hasPrev := set[v]
		set[v] = i
		if prev < startAt {
			// 現在の候補が確定
			interval = i - (startAt - 1)
		} else if hasPrev {
			// startAt と prev または prev と i の間にある二つの候補から確定
			interval = max(prev-startAt, i-prev)
			// 候補のスタート位置を更新
			startAt = prev + 1
		} else {
			interval = i - (prev - 1)
		}
		// 確定した候補の文字列が長さを更新するか
		if interval > maxCounts {
			maxCounts = interval
		}
	}

	return maxCounts
}
