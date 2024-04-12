package easy

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	// create a string builder and ensure it has enough capacity
	var sb strings.Builder
	sb.Grow(len(word1) + len(word2))

	// merge alternatively till the end of the shorter word
	i, j := 0, 0
	for i < len(word1) && j < len(word2) {
		sb.WriteByte(word1[i])
		sb.WriteByte(word2[j])
		i++
		j++
	}

	// merge the rest of word1 if any
	for i < len(word1) {
		sb.WriteByte(word1[i])
		i++
	}

	// merge the rest of word2 if any
	for j < len(word2) {
		sb.WriteByte(word2[j])
		j++
	}

	// return the merged string
	return sb.String()
}
