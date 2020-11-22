package leetcode

func isPalindromic(str string) bool {
	l := len(str)
	if l <= 1 {
		return true
	}
	ok := true
	for i := 0; i < l/2; i++ {
		ok = ok && (str[i] == str[l-i-1])
	}
	return ok
}

func longestPalindrome(s string) string {
	l := len(s)
	var longest string
	for i := 0; i <= l; i++ {
		for j := i + 1; j <= l; j++ {
			str := s[i:j]
			if len(str) > len(longest) && isPalindromic(str) {
				longest = str
			}
		}
	}
	return longest
}
