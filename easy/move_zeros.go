package easy

func moveZeroes(nums []int) {
	i, j := 0, 0
	for j < len(nums) {
		// if current right is 0, go next and check
		if nums[j] == 0 {
			j++
			continue
		}
		// swap whenever right is not 0 and the two element is not the same
		if i != j {
			nums[i], nums[j] = nums[j], nums[i]
		}
		// move to the next set
		i++
		j++
	}
}
