package medium

import (
	"sort"
)

func deckRevealedIncreasing(deck []int) []int {
	// sort the deck in descending order, so that we can work backwards
	sort.Sort(sort.Reverse(sort.IntSlice(deck)))

	// create a new slice to store the result
	// double the length to prevent expanding the slice later
	output := make([]int, len(deck)*2)
	totalLen := len(output) - 1
	output[totalLen] = deck[0]
	start, end := totalLen-1, totalLen

	// iterate through the deck and fill the output slice
	for i := 1; i < len(deck); i++ {
		// working backwards, we move the last element to the front (second place)
		output[start] = output[end]
		// assign the current element with the revealed card
		output[start-1] = deck[i]
		// since we filled two elements, we move the start pointer back by 2
		start -= 2
		// the end is now duplicated, so we move it back by 1
		end--
	}
	return output[start+1 : end+1]
}
