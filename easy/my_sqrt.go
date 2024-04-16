package easy

func mySqrt(x int) int {
	// use binary search to solve the problem
	// the possible n in n to the power of 2
	left := 0
	right := x
	for left <= right {
		// middle value
		mid := (left + right) / 2

		// found the vlaue
		if mid*mid == x {
			return mid
			// too small, elevate lower bound
		} else if mid*mid < x {
			left = mid + 1
			// too big,
		} else {
			right = mid - 1
		}
	}
	// round value, for no n
	return left - 1
}
