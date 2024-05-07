package easy

func uniqueOccurrences(arr []int) bool {
	// Create a map to record the number of times each integer occurs
	freqCount := make(map[int]int)
	for _, v := range arr {
		_, existed := freqCount[v]
		// if not recorded yet, init the count to 1
		if !existed {
			freqCount[v] = 1
			continue
		}
		// otherwise, increment the count by 1
		freqCount[v]++
	}

	// Create another map to record check whether the frequency count of each integer is unique
	uniqueNum := make(map[int]bool)
	for _, v := range freqCount {
		_, existed := uniqueNum[v]
		// if the frequency count exists, return false
		if existed {
			return false
		} else {
			// else record it
			uniqueNum[v] = true
		}
	}
	return true
}
