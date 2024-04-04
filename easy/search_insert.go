package easy

func searchInsert(nums []int, target int) int {
	// just a binary search
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// return left because left is the index where target should be inserted
	return left
}
