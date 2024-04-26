package easy

func sumIntSlice(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func pivotIndex(nums []int) int {
	// left start at 0, and right start at the sum of the slice
	left, right := 0, sumIntSlice(nums)

	for i, v := range nums {
		// if the sum of v's left side is equal to the sum of v's right side, return the index
		if left == right-left-v {
			return i
		}
		// accumulate the left side
		left += v
	}
	// no pivot index found
	return -1
}
