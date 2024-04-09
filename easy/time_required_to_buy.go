package easy

func timeRequiredToBuy(tickets []int, k int) int {
	i := 0
	seconds := 0
	for true {
		// if target has no ticket to buy, break
		if tickets[k] == 0 {
			break
		}

		// if i is over the length, back to start
		if i >= len(tickets) {
			i = 0
		}

		// if the current value has no ticket to buy, skip it
		if tickets[i] == 0 {
			i++
			continue
		}

		// buy ticket, move to next, increase total seconds
		tickets[i]--
		i++
		seconds++
	}
	return seconds
}
