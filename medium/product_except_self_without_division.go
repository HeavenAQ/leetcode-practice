package medium

func productExceptSelf(nums []int) []int {
	n := len(nums)

	// calculate the products to the left of the current element
	pre := make([]int, n)
	pre[0] = 1
	for i := 0; i < n-1; i++ {
		pre[i+1] = pre[i] * nums[i]
	}

	// calculate the products to the right of the current element
	suf := make([]int, n)
	suf[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		suf[i] = suf[i+1] * nums[i+1]
	}

	// store the answers
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = pre[i] * suf[i]
	}
	return ans
}
