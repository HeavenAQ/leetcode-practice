package medium

import (
	"strconv"
)

/* func: storeIntToByte
*  convert the current count to byte and stor it to the chars slice
*  @param chars: the input slice
*  @param cur: the index indicating where to store the character count to the chars
*  @param count: the count of consecutively repeated characters
*
*  @return: return the next position to store the character count
 */
func storeIntToByte(chars []byte, cur int, count int) int {
	intString := strconv.Itoa(count)
	for _, v := range intString {
		chars[cur] = byte(v)
		cur++
	}
	return cur
}

func compress(chars []byte) int {
	n := len(chars)
	// if only one character in the slice, return
	if n == 1 {
		return 1
	}

	// cur indexes and helps overwrite the value stored in chars
	// count records each character count
	cur, count := 0, 1

	// start from the second character, and compare the current character with the previous one in each round
	for i := 1; i < n; i++ {
		// if the current character is the same as the last one, increase count
		if chars[i] == chars[i-1] {
			count++
			// if the current character stops repeating
		} else {
			// record the character
			chars[cur] = chars[i-1]
			cur++
			// store the count and move cur to the next position
			if count > 1 {
				cur = storeIntToByte(chars, cur, count)
			}
			// reset the count
			count = 1
		}
	}

	// ensure the last character is recorded
	chars[cur] = chars[n-1]
	cur++

	// if the last character in chars is repeated, record its count
	if count > 1 {
		cur = storeIntToByte(chars, cur, count)
	}

	// return the length used to record the characters and their counts
	return cur
}
