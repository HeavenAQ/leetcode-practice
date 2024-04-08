package medium

func reverseWords(s string) string {
	reversed := ""

	// iterate through the string and reverse each word
	for i := 0; i < len(s); {
		// if the character is not a space, then it is a character of a word
		if s[i] != ' ' {
			// empty string to store the word
			word := ""
			for true {
				if i == len(s) || s[i] == ' ' {
					break
				}
				// concatenate the current character and move to the next character
				word += string(s[i])
				i++
			}
			// prepend the word to the reversed string
			reversed = word + " " + reversed
		}

		i++
	}
	return reversed[:len(reversed)-1]
}
