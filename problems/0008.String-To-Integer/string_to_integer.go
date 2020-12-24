package leetcode

func myAtoi(s string) int {
	const (
		MaxInt = int(^uint32(0) >> 1)
		MinInt = -MaxInt - 1
	)
	var (
		num  int
		sign int8
	)

	for _, v := range s {
		switch {
		case sign == 0 && v == ' ':
			continue
		case sign == 0 && v == '+':
			sign = 1
		case sign == 0 && v == '-':
			sign = -1
		case '0' <= v && v <= '9':
			if sign == 0 {
				sign = 1
			}
			num = num*10 + int(int8(v-'0')*sign)
			if num < MinInt || MaxInt < num {
				if sign < 0 {
					num = MinInt
				} else {
					num = MaxInt
				}
				return num
			}
		default:
			return num
		}
	}

	return num
}
