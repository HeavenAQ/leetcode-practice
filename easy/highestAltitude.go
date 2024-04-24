package easy

import "math"

func largestAltitude(gain []int) int {
	max := int(math.Inf(-1))
	start := 0
	for _, v := range gain {
		if max < start {
			max = start
		}
		start += v
	}

	if max < start {
		return start
	}
	return max
}
