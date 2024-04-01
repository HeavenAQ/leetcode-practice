package medium

import "strings"

func lengthOfLongestSubstring(s string) int {
	longestLen := 0  // longest length of substring
	currentStr := "" // longest substring
	for _, char := range s {
		/*
			Purpose:
			if current character is in the longest substring, remove the characters before it, and add
			the current character to the current substring

			Why:
			- Since we want to find the longest substring (not sub-sequence) without repeating characters,
			we need to remove characters from the start of the current substring to the first occurrence
			of our current character. After that,
			- e.g. "pwwkew"
				* currentStr = "pw"
				* char = 'w'
				* 1st iteration: currentStr = "w"
				* 2nd iteration: currentStr = ""
				* currentStr += string(char): "w"
		*/
		for strings.Contains(currentStr, string(char)) {
			currentStr = currentStr[1:]
		}
		currentStr += string(char)

		/*
			Purpose:
			if the current substring is longer than the longest substring, update the longest substring length
		*/
		currentStrLen := len(currentStr)
		if longestLen < currentStrLen {
			longestLen = currentStrLen
		}
	}
	return longestLen
}
