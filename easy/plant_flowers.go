package easy

func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	for i := 0; i < length; i++ {
		// if the place is taken, go to next
		if flowerbed[i] == 1 {
			continue
		}

		// check current left and right are both emtpy
		left := (i == 0 || flowerbed[i-1] == 0)
		right := (i == length-1 || flowerbed[i+1] == 0)

		// if both empty, plant it
		if left && right {
			flowerbed[i] = 1
			n--
		}
	}
	return n <= 0
}
