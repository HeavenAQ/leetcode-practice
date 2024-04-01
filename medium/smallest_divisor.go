package medium

import "math"

func getSum(nums []int, divisor int) int {
	/*
		Purpose:
		Divide every number in nums by the divisor and take the ceiling of the outcomes to sum them up
	*/
	sum := 0
	for _, v := range nums {
		sum += int(math.Ceil(float64(v) / float64(divisor)))
	}
	return sum
}

func max(nums []int) int {
	/*
		Purpose:
		Get the max value from the slice
	*/
	maxVal := 0
	for _, v := range nums {
		if maxVal < v {
			maxVal = v
		}
	}
	return maxVal
}

func smallestDivisor(nums []int, threshold int) int {
	/*
		Purpose:
		Use binary search to get the divisor
	*/
	left := 1          // lower bound
	right := max(nums) // upper bound

	// while lower bound is not greater than upper bound
	for left <= right {
		// get the middle of the upper and lower bound
		// left -> offset
		// (right - left) / 2 -> middle
		mid := left + (right-left)/2

		// if the sum is smaller than the threshold
		if getSum(nums, mid) <= threshold {
			// The divisor is too large. Get a smaller divisor to have a bigger sum.
			right = mid - 1
		} else {
			//  The divisor is too small. Get A larger divisor to have a smaller sum.
			left = mid + 1
		}
	}

	// if left == right -> found the divisor that helps create a sum == threshold
	// if left > right -> found the divisor that helps create a sum < threshold
	return left
}
