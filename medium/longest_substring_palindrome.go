package medium

import "strings"

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

// Manacher's Algorithm
// In the Manacher's Algorithm, we preprocess the string to insert '#' between each character
func preprocessStr(s string) string {
	var sb strings.Builder
	sb.Grow(len(s)*2 + 1)
	sb.WriteByte('#')
	for _, v := range s {
		sb.WriteByte(byte(v))
		sb.WriteByte('#')
	}
	return sb.String()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func manacherLongestSubPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	maxLen := 1
	maxStr := string(s[0])
	s = preprocessStr(s)
	palindromeRadii := make([]int, len(s)) // palindromeRadii[i] = the radius of the palindrome centered at i
	center, right := 0, 0

	for i := range s {
		// The radius of the palindrome centered at i is at least 1
		// If i < right, then the radius of the palindrome centered at i is at least min(right-i, palindromeRadii[2*center-i])
		// right - i: the distance from i to right
		// palindromeRadii[2*center-i]: the radius of the palindrome centered at 2*center-i (mirrored index)
		if i < right {
			palindromeRadii[i] = min(right-i, palindromeRadii[2*center-i])
		}

		// Try to expand palindrome centered at i
		radius := palindromeRadii[i]
		for i-radius-1 >= 0 && i+radius+1 < len(s) && s[i-radius-1] == s[i+radius+1] {
			palindromeRadii[i]++
			radius++
		}

		// Update center and right if the palindrome centered at i expands past right
		if i+radius > right {
			center, right = i, i+radius
		}

		// Update the maximum length and the corresponding string
		if radius > maxLen {
			maxLen = radius
			maxStr = s[i-radius : i+radius+1]
		}
	}

	// Remove '#' from the final string
	return strings.Replace(maxStr, "#", "", -1)
}
