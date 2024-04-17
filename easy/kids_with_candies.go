package easy

// function for finding the max value in the slice
func max(a []int) int {
	max := 0
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxVal := max(candies)
	results := make([]bool, len(candies))
	for i, v := range candies {
		// if the current value with extraCandies could be larger than or equal to result
		if v+extraCandies >= maxVal {
			results[i] = true
		} else {
			results[i] = false
		}
	}
	return results
}
