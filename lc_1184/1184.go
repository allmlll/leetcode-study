package lc_1184

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	sum := 0
	L := 0
	for i := 0; i < len(distance); i++ {
		sum += distance[i]
		if i >= min(start, destination) && i <= max(start, destination) {
			L += distance[i]
		}

	}
	return min(sum-L, L)
}
