package easy

// Given that we take either 1 or 2 step each round until we reach n
// That is either of the following:
// 1. 1 stair from the (n-1)th stair
// 2. 2 stair from the (n-2)th stair
// It's a fibonacci sequence.
// Space Optimzed Solution (iterative fibonacci)
func climbStairs(n int) int {
	if n == 1 {
		return n
	}

	n1 := 1
	n2 := 2
	for i := 3; i <= n; i++ {
		prev := n1     // value of i - 2
		n1 = n2        // value of i - 1
		n2 = prev + n1 // value of i
	}
	return n2
}
