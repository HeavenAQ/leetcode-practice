package easy

import "strings"

func reversePrefix(word string, ch byte) string {
	// find ch in the word
	for i, v := range word {
		// found ch
		if byte(v) == ch {
			// string builder for efficient string manipulation
			var sb strings.Builder
			sb.Grow(len(word))

			// write the characters before ch in reverse order
			for j := i; j >= 0; j-- {
				sb.WriteByte(word[j])
			}

			// write the characters after ch
			for j := i + 1; j < len(word); j++ {
				sb.WriteByte(word[j])
			}

			// return new word
			return sb.String()
		}
	}

	// if no ch, return word
	return word
}
