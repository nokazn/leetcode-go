package leetcode

const digit = 10

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var nums []int
	for {
		nums = append(nums, x%digit)
		x /= digit
		if x < 1 {
			break
		}
	}

	output := true
	l := len(nums)
	for i := 0; i < l/2; i++ {
		output = output && (nums[i] == nums[l-i-1])
		if !output {
			break
		}
	}
	return output
}

func isPalindrome2(x int) bool {
	if x < 0 || (x%10 == 0 && x > 0) {
		return false
	}

	var reverted int
	for x > reverted {
		reverted = reverted*digit + x%digit
		x /= digit
	}
	return x == reverted || x == reverted/10
}
