package medium

func LongestSubstringLength(s string) int {
	if len(s) == 0 {
		return 0
	}

	longestLen := 0
	curStartIdx := 0
	appeared := make(map[rune]int)

	for i, v := range s {
		if prevIdx, exists := appeared[v]; exists && prevIdx >= curStartIdx {
			// If the character exists and is within the current substring,
			// move the start right after the previous occurrence of `v`
			curStartIdx = prevIdx + 1
		}

		// Update the character's latest index
		appeared[v] = i

		// Calculate the current length and update longestLen if necessary
		curLen := i - curStartIdx + 1
		if curLen > longestLen {
			longestLen = curLen
		}
	}
	return longestLen
}
