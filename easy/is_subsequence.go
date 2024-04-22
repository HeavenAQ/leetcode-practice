package easy

func isSubsequence(s string, t string) bool {
	i := 0
	for _, v := range t {
		// if all matched, retur true
		if i == len(s) {
			return true
			// if the current char in t equals the the current char in s, increment s's index
		} else if rune(s[i]) == v {
			i++
		}
	}

	// if nothing left to check in s, return true
	if i == len(s) {
		return true
	}
	// else return false
	return false
}
