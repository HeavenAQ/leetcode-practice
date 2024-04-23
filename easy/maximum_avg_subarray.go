package easy

func sum(nums []int, k int) float64 {
	total := 0.0
	for i := 0; i < k; i++ {
		total += float64(nums[i])
	}
	return total
}

func findMaxAverage(nums []int, k int) float64 {
	// edge case: if k > the length of nums, return
	if len(nums) < k {
		return 0
	}

	currentSum := sum(nums, k) // record the sliding window
	maxSum := currentSum       // record the max
	for i := 1; i <= len(nums)-k; i++ {
		// Instead of calculating the average every time, we use sum to check
		// if the current value is larger than the max
		currentSum := currentSum - float64(nums[i-1]) + float64(nums[i+k-1])
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum / float64(k)
}
