package easy

import (
	"slices"
)

func reverseVowels(s string) string {
	out := []rune(s)
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	i, j := 0, len(s)-1

	// two idices moving in different directions
	for i <= j {
		// if both current runes are vowels, swap
		if slices.Contains(vowels, out[i]) && slices.Contains(vowels, out[j]) {
			tmp := out[i]
			out[i] = out[j]
			out[j] = tmp
			i++
			j--
			// if ith rune is vowel but jth rune is not, move j to next
		} else if slices.Contains(vowels, out[i]) {
			j--
			// if jth rune is vowel but ith rune is not, move i to next
		} else if slices.Contains(vowels, out[j]) {
			i++
			// if both runes are consonant, move to next
		} else {
			i++
			j--
		}
	}
	return string(out)
}
