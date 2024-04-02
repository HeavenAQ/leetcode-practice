package medium

import "strings"

func subPalindrome(s string, output *string, left, right int) {
	strLen := len(s)
	// spread to the left and right and check if the characters on the left and right are the same
	for left >= 0 && right < strLen {
		prev := s[left]
		next := s[right]
		if prev == next {
			*output = string(prev) + *output + string(next)
		} else {
			break
		}
		left--
		right++
	}
}

func maxStrLen(a, b string) string {
	if len(a) > len(b) {
		return a
	} else {
		return b
	}
}

func longestPalindrome(s string) string {
	// recording the current longest
	longestStr := ""
	for i, v := range s {
		// odd number
		// initialize the current string with the current character
		// e.g.
		// "aba"
		// oddCurStr = "b"
		// left = 0, right = 2
		// prev = a, next = a
		oddCurStr := string(v)
		subPalindrome(s, &oddCurStr, i-1, i+1)

		// even number
		// initialize the current string with ""
		// e.g.
		// "abba"
		// evenCurStr = ""
		// left = 1, right = 2
		// 1st round: prev = b, next = b
		// 2nd round: prev = a, next = a
		evenCurStr := ""
		subPalindrome(s, &evenCurStr, i, i+1)

		// update the longest string
		longerStr := maxStrLen(evenCurStr, oddCurStr)
		longestStr = strings.Clone(maxStrLen(longerStr, longestStr))
	}
	return longestStr
}
