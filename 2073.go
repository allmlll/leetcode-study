package main

func timeRequiredToBuy(tickets []int, k int) int {
	res := 0
	for _, ticket := range tickets {
		if ticket > tickets[k] {
			res += tickets[k]
			ticket -= tickets[k]
		} else {
			res += ticket
			ticket = 0
		}

	}
	return res

}
