package medium

func subPalindrome(s string, left, right int) string {
	// spread to the left and right and check if the characters on the left and right are the same
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

func maxStr(a, b string) string {
	if len(a) > len(b) {
		return a
	} else {
		return b
	}
}

func longestPalindrome(s string) string {
	// recording the current longest
	longestStr := ""
	for i := range s {
		// odd number
		// initialize the current string with the current character
		// e.g.
		// "aba"
		// oddCurStr = "b"
		// left = 0, right = 2
		// prev = a, next = a
		oddCurStr := subPalindrome(s, i, i)

		// even number
		// initialize the current string with ""
		// e.g.
		// "abba"
		// evenCurStr = ""
		// left = 1, right = 2
		// 1st round: prev = b, next = b
		// 2nd round: prev = a, next = a
		evenCurStr := subPalindrome(s, i, i+1)

		// update the longest string
		longestStr = maxStr(longestStr, maxStr(evenCurStr, oddCurStr))
	}
	return longestStr
}
