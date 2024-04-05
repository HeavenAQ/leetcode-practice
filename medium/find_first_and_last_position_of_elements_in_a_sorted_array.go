package medium

func searchRange(nums []int, target int) []int {
	// implement binary search
	// empty array, return [-1, -1]
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	// setup variables
	max, min := -1, -1
	left, right := 0, len(nums)-1

	// if left > right, not found -> exit
	for left <= right {
		// find the middle index
		mid := (left + right) / 2
		// if found, find the min and max index
		if nums[mid] == target {
			// find the min
			for i := 0; mid-i >= 0 && nums[mid-i] == target; i++ {
				min = mid - i
			}

			// find the max
			for i := 0; mid+i < len(nums) && nums[mid+i] == target; i++ {
				max = mid + i
			}
			break
			// too big, move left
		} else if nums[mid] > target {
			right = mid - 1
			// too small, move right
		} else {
			left = mid + 1
		}
	}
	return []int{min, max}
}
