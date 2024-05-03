package easy

func findDifference(nums1 []int, nums2 []int) [][]int {
	// to maps and use them like sets
	nums1Map := make(map[int]bool)
	nums2Map := make(map[int]bool)

	// store the numbers in both slices as keys
	for _, v := range nums1 {
		nums1Map[v] = false
	}
	for _, v := range nums2 {
		nums2Map[v] = false
	}

	// get the numbers that do not exist in the other array
	res := [][]int{{}, {}}
	for k := range nums1Map {
		_, existed := nums2Map[k]
		if !existed {
			res[0] = append(res[0], k)
		}
	}

	for k := range nums2Map {
		_, existed := nums1Map[k]
		if !existed {
			res[1] = append(res[1], k)
		}
	}

	return res
}
